// +build ignore

/*
 * Copyright (C) 2019 The "MysteriumNetwork/go-dvpn-web" Authors.
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
