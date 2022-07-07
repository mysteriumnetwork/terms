// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package generate

import (
	"strings"

	"github.com/mysteriumnetwork/go-ci/env"
	"github.com/mysteriumnetwork/terms/ci/buildenv"
)

func NextVersionUnPrefixed() string {
	version := env.Str(buildenv.NextVersion)

	if strings.HasPrefix(version, "v") {
		return version[1:]
	}

	return version
}
