// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

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

	_, err = git.Commit("Update terms packages [skip ci]",
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
