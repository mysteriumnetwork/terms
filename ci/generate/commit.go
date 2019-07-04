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
	cienv "github.com/mysteriumnetwork/go-ci/env"
	mgit "github.com/mysteriumnetwork/go-ci/git"
	env "github.com/mysteriumnetwork/terms/ci/env"
	"github.com/pkg/errors"
)

// Commits and pushes generated files
func CommitGenerated() error {
	err := cienv.EnsureEnvVars(cienv.GithubOwner, cienv.GithubRepository, cienv.GithubAPIToken, env.NextVersion)
	if err != nil {
		return err
	}

	git := mgit.NewCommiter(cienv.Str(cienv.GithubAPIToken))

	err = git.Checkout(&mgit.CheckoutOptions{
		BranchName: "master",
		Keep:       true,
	})
	if err != nil {
		return err
	}

	_, err = git.Commit("Update terms packages",
		"terms-go/terms_bindata.go",
		"terms-js/package.json",
		"terms-js/index.js",
	)
	if err != nil {
		return errors.Wrap(err, "could not commit generated files")
	}
	err = git.Push(&mgit.PushOptions{
		Remote: "upstream",
	})
	if err != nil {
		return errors.Wrap(err, "could not push to upstream")
	}

	return nil
}
