// +build mage

/*
 * Copyright (C) 2019 The "MysteriumNetwork/node" Authors.
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

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mysteriumnetwork/terms/terms-go"
)

// Generate generates packages for all language targets
func Generate() error {
	mg.SerialDeps(
		GenerateGo,
		GenerateJs,
	)
	return nil
}

// GenerateGo embeds terms into `terms-go/terms-bindata.go`
func GenerateGo() error {
	return sh.RunV("go", "generate", "./...")
}

type templateParams struct {
	BuildVersion string
	TermsMd      string
	WarrantyMd   string
}

// GenerateJs embeds terms into `terms-js/index.js`
func GenerateJs() error {
	params := &templateParams{
		BuildVersion: os.Getenv("BUILD_VERSION"),
		TermsMd:      generateJsMultiline(string(terms.TermsMdBytes)),
		WarrantyMd:   generateJsMultiline(string(terms.WarrantyMdBytes)),
	}

	templateFiles, err := ioutil.ReadDir("terms-js/template")
	if err != nil {
		return err
	}

	for _, f := range templateFiles {
		if err := generateJsUsingTemplate(filepath.Join("terms-js/template", f.Name()), filepath.Join("terms-js", f.Name()), params); err != nil {
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

	template, err := template.New(filepath.Base(templatePath)).
		Funcs(template.FuncMap{"noescape": noescape}).
		ParseFiles(templatePath)
	if err != nil {
		return err
	}

	return template.Execute(targetFile, params)
}
