// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package generate

import (
	cienv "github.com/mysteriumnetwork/go-ci/env"
	"github.com/mysteriumnetwork/terms/ci/env"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var link = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

func toCamelCase(str string) string {
	return link.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

func FileNameToVariableName(name string) string {
	return toCamelCase(strings.Replace(strings.ToLower(name), ".md", "", 1))
}

func GetDocumentPaths(dir string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func NextVersionUnPrefixed() string {
	version := cienv.Str(env.NextVersion)

	if strings.HasPrefix(version, "v") {
		return version[1:]
	}

	return version
}
