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

package env

import (
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
