package publish

import (
	"strconv"
	"strings"

	"github.com/mysteriumnetwork/go-ci/env"
	cigit "github.com/mysteriumnetwork/go-ci/git"
	"github.com/mysteriumnetwork/go-ci/github"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// Publish commits & pushes changes
func Publish() error {
	must(env.EnsureEnvVars(env.GithubAPIToken))
	git := cigit.NewCommiter(env.Str(env.GithubAPIToken))

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
	releaser, err := github.NewReleaser(env.Str(env.GithubOwner), env.Str(env.GithubRepository), env.Str(env.GithubAPIToken))
	if err != nil {
		return err
	}

	latest, err := releaser.Latest()
	if err != nil {
		return err
	}
	nextTag, err := nextTag(latest.TagName)
	if err != nil {
		return err
	}

	if _, err := releaser.Create(nextTag); err != nil {
		return err
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
