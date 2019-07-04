package publish

import (
	"strconv"
	"strings"

	"github.com/mysteriumnetwork/go-ci/env"
	cigit "github.com/mysteriumnetwork/go-ci/git"
	"github.com/mysteriumnetwork/go-ci/github"
	"github.com/pkg/errors"
)

// Publish commits & pushes changes
func Publish() error {
	err := env.EnsureEnvVars(env.GithubAPIToken)
	if err != nil {
		return err
	}

	git := cigit.NewCommiter(env.Str(env.GithubAPIToken))
	err = git.Checkout("master")
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
	err = git.Push(&cigit.PushOptions{
		Remote: "upstream",
	})
	if err != nil {
		return errors.Wrap(err, "could not push to upstream")
	}
	releaser, err := github.NewReleaser(env.Str(env.GithubOwner), env.Str(env.GithubRepository), env.Str(env.GithubAPIToken))
	if err != nil {
		return errors.Wrap(err, "could not create github releaser")
	}

	latest, err := releaser.Latest()
	if err != nil {
		return errors.Wrap(err, "could not get latest release")
	}
	nextTag, err := nextTag(latest.TagName)
	if err != nil {
		return errors.Wrap(err, "could not calculate next tag")
	}

	if _, err := releaser.Create(nextTag); err != nil {
		return errors.Wrap(err, "could not create next tag")
	}

	return nil
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
