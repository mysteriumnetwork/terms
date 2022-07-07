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

	"github.com/mysteriumnetwork/go-ci/env"
	"github.com/mysteriumnetwork/terms/ci/buildenv"
	"github.com/mysteriumnetwork/terms/documents"
)

var documentVars = map[string]string{
	"TermsBountyPilot": generateJsMultiline(documents.TermsBountyPilot),
	"TermsEndUser":     generateJsMultiline(documents.TermsEndUser),
	"TermsExitNode":    generateJsMultiline(documents.TermsExitNode),
	"TermsNodeShort":   generateJsMultiline(documents.TermsNodeShort),
	"Warranty":         generateJsMultiline(documents.Warranty),
}

type templateParams struct {
	BuildVersion string
	UpdatedAt    string
	Documents    map[string]string
}

// GenerateJs embeds terms into `terms-js/index.js`
func GenerateJs() error {
	err := env.EnsureEnvVars(buildenv.NextVersion)
	if err != nil {
		return err
	}

	params := &templateParams{
		BuildVersion: NextVersionUnPrefixed(),
		UpdatedAt:    time.Now().Format("2006-01-02"),
		Documents:    documentVars,
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
