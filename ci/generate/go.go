// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

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
