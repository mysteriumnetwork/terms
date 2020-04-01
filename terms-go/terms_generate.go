// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

// +build ignore

package main

import (
	"log"

	"github.com/go-bindata/go-bindata"

	"github.com/mysteriumnetwork/terms/ci/generate"
)

func main() {
	cfg := &bindata.Config{
		Package: "terms",
		Input: []bindata.InputConfig{
			{Path: "../documents/"},
		},
		Output: "./terms_bindata.go",
		Prefix: "../documents/",
	}
	err := bindata.Translate(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	err = generate.GenerateGoVariables()
	if err != nil {
		log.Fatalln(err)
	}
}
