// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package env

import (
	"fmt"
	"strconv"
	"strings"

	cienv "github.com/mysteriumnetwork/go-ci/env"
	"github.com/mysteriumnetwork/go-ci/github"
	"github.com/pkg/errors"
)

const NextVersion = cienv.BuildVar("NEXT_VERSION")

// GenerateEnvFile for sourcing in other stages
func GenerateEnvFile() error {
	err := cienv.EnsureEnvVars(cienv.GithubOwner, cienv.GithubRepository, cienv.GithubAPIToken)
	if err != nil {
		return err
	}

	next, err := NextTag()
	if err != nil {
		return err
	}

	vars := []cienv.EnvVar{
		{NextVersion, next},
	}
	return cienv.WriteEnvVars(vars, "./build/env.sh")
}

// NextTag determines next tag in repo
func NextTag() (string, error) {
	releaser, err := github.NewReleaser(cienv.Str(cienv.GithubOwner), cienv.Str(cienv.GithubRepository), cienv.Str(cienv.GithubAPIToken))
	if err != nil {
		return "", errors.Wrap(err, "could not create github releaser")
	}

	latest, err := releaser.Latest()
	if err != nil {
		return "", errors.Wrap(err, "could not get latest release")
	}

	next, err := nextTag(latest.TagName)
	if err != nil {
		return "", errors.Wrap(err, "could not calculate next tag")
	}

	if !strings.HasPrefix(next, "v") {
		return fmt.Sprintf("v%s", next), nil
	}

	return next, nil
}

func nextTag(previous string) (string, error) {
	parts := strings.Split(previous, ".")
	patch, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return "", err
	}
	patch = patch + 1
	parts[len(parts)-1] = strconv.Itoa(patch)
	return strings.Join(parts, "."), nil
}
