package generate

import (
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
