// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

// +build mage

package main

import (
	// mage:import
	_ "github.com/mysteriumnetwork/terms/ci/env"
	// mage:import
	_ "github.com/mysteriumnetwork/terms/ci/generate"
	// mage:import
	_ "github.com/mysteriumnetwork/terms/ci/publish"
)
