// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package publish

import (
	cienv "github.com/mysteriumnetwork/go-ci/env"
	"github.com/mysteriumnetwork/go-ci/github"
	"github.com/mysteriumnetwork/terms/ci/buildenv"
	"github.com/pkg/errors"
)

// PublishGithub publishes next Github release
func PublishGithub() error {
	nextVersion := cienv.Str(buildenv.NextVersion)

	releaser, err := github.NewReleaser(cienv.Str(cienv.GithubOwner), cienv.Str(cienv.GithubRepository), cienv.Str(cienv.GithubAPIToken))
	if err != nil {
		return err
	}

	if _, err := releaser.Create(nextVersion); err != nil {
		return errors.Wrap(err, "could not create next tag")
	}

	return nil
}
