// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package generate

import (
	"io/ioutil"

	"github.com/magefile/mage/sh"
)

const TermsVersionVarName = "TermsVersion"

func GenerateGoVersion() error {
	vars := "\n\t" + TermsVersionVarName + " = " + `"` + NextVersionUnPrefixed() + `"`
	str := `package terms

var (` + vars + `
)`

	err := ioutil.WriteFile("./../terms-go/version.go", []byte(str), 0644)
	if err != nil {
		return err
	}

	return nil
}

// GenerateGo embeds terms into `terms-go`
func GenerateGo() error {
	return sh.RunV("go", "generate", "./...")
}
