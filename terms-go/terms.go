// Copyright (c) 2020 BlockDev AG
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package terms

import (
	"github.com/mysteriumnetwork/terms/documents"
)

//go:generate go run terms_generate.go

var (
	TermsBountyPilot = documents.TermsBountyPilot
	TermsEndUser     = documents.TermsEndUser
	TermsExitNode    = documents.TermsExitNode
	TermsNodeShort   = documents.TermsNodeShort
	Warranty         = documents.Warranty
)
