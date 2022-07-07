// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

//go:build ignore

package main

import (
	"log"

	"github.com/mysteriumnetwork/terms/ci/generate"
)

func main() {
	if err := generate.GenerateGoVersion(); err != nil {
		log.Fatalln(err)
	}
}
