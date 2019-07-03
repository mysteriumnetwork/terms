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
	"github.com/mysteriumnetwork/go-ci/env"
	cigit "github.com/mysteriumnetwork/go-ci/git"
	"github.com/mysteriumnetwork/terms/ci/generate"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// CI generates, commits, releases & pushes updated terms
func CI() error {
	must(env.EnsureEnvVars(env.GithubAPIToken))
	git := cigit.NewCommiter(env.Str(env.GithubAPIToken))
	must(git.Checkout("master"))
	must(generate.Generate())
	_, err := git.Commit("Updating terms packages",
		"terms-go/terms_bindata.go",
		"terms-js/package.json",
		"terms-js/index.js",
	)
	if err != nil {
		return err
	}
	must(git.Push(&cigit.PushOptions{
		Remote: "upstream",
	}))
	return nil
}
