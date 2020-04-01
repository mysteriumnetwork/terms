/*
 * Copyright (C) 2019 The "MysteriumNetwork/terms" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
		BuildVersion: cienv.Str(env.NextVersion),
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
