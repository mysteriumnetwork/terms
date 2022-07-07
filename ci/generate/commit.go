// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package generate

import (
	"github.com/mysteriumnetwork/go-ci/env"
	mgit "github.com/mysteriumnetwork/go-ci/git"
	"github.com/mysteriumnetwork/terms/ci/buildenv"
	"github.com/pkg/errors"
)

// CommitGenerated commits and pushes generated files.
func CommitGenerated() error {
	err := env.EnsureEnvVars(env.GithubOwner, env.GithubRepository, env.GithubAPIToken, buildenv.NextVersion)
	if err != nil {
		return err
	}

	git := mgit.NewCommiter(env.Str(env.GithubAPIToken))

	err = git.Checkout(&mgit.CheckoutOptions{
		BranchName: "master",
		Keep:       true,
	})
	if err != nil {
		return err
	}

	_, err = git.Commit("Update terms packages [skip ci]",
		"terms-go/version.go",
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
