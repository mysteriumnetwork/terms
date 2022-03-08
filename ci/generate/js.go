// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package generate

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	cienv "github.com/mysteriumnetwork/go-ci/env"

	"github.com/mysteriumnetwork/terms/ci/env"
	"github.com/mysteriumnetwork/terms/terms-go"
)

type templateVariables map[string]string
type templateParams struct {
	BuildVersion string
	UpdatedAt    string
	Documents    templateVariables
}

func mapDocumentsIntoVariables(files []os.FileInfo) (templateVariables, error) {
	variables := make(templateVariables)

	for _, file := range files {
		key := FileNameToVariableName(file.Name())
		data, err := ioutil.ReadFile(filepath.Join(terms.DocumentDirectory, file.Name()))
		if err != nil {
			return nil, err
		}

		variables[key] = generateJsMultiline(string(data))
	}

	return variables, nil
}

// GenerateJs embeds terms into `terms-js/index.js`
func GenerateJs() error {
	err := cienv.EnsureEnvVars(env.NextVersion)
	if err != nil {
		return err
	}

	documents, err := GetDocumentPaths(terms.DocumentDirectory + "/")
	if err != nil {
		return err
	}

	variables, err := mapDocumentsIntoVariables(documents)
	if err != nil {
		return err
	}

	params := &templateParams{
		BuildVersion: nextJSVersion(),
		UpdatedAt:    time.Now().Format("2006-01-02"),
		Documents:    variables,
	}

	templateFiles, err := ioutil.ReadDir("terms-js/template")
	if err != nil {
		return err
	}

	for _, f := range templateFiles {
		err := generateJsUsingTemplate(
			filepath.Join("terms-js/template", f.Name()),
			filepath.Join("terms-js", f.Name()),
			params,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func nextJSVersion() string {
	version := cienv.Str(env.NextVersion)
	if strings.HasPrefix(version, "v") {
		return version[1:len(version)]
	}
	return version
}

func generateJsMultiline(in string) string {
	lines := strings.Split(in, "\n")
	for i, line := range lines {
		line = template.JSEscapeString(line)
		line = fmt.Sprintf("'%s \\n'", line)
		lines[i] = line
	}

	multiline := strings.Join(lines, " + \n")
	return multiline
}

func noescape(s string) template.HTML {
	return template.HTML(s)
}

func generateJsUsingTemplate(templatePath, targetPath string, params *templateParams) error {
	targetFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	tpl, err := template.New(filepath.Base(templatePath)).
		Funcs(template.FuncMap{"noescape": noescape}).
		ParseFiles(templatePath)
	if err != nil {
		return err
	}

	return tpl.Execute(targetFile, params)
}
