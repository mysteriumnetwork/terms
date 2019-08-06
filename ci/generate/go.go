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
	"io/ioutil"

	"github.com/magefile/mage/sh"

	"github.com/mysteriumnetwork/terms/terms-go"
)

func GenerateGoVariables() error {
	docs, err := GetDocumentPaths("../" + terms.DocumentDirectory)
	if err != nil {
		return err
	}

	vars := ""
	for _, doc := range docs {
		vars += "\n\t" + FileNameToVariableName(doc.Name()) + ` = MustAsset("` + doc.Name() + `")`
	}
	str := `package terms

var (` + vars + `
)`

	err = ioutil.WriteFile("./../terms-go/variables.go", []byte(str), 0644)
	if err != nil {
		return err
	}

	return nil
}

// GenerateGo embeds terms into `terms-go/terms-bindata.go`
func GenerateGo() error {
	return sh.RunV("go", "generate", "./...")
}
