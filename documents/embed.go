package documents

import (
	_ "embed"
)

var (
	//go:embed TERMS_BOUNTY_PILOT.md
	TermsBountyPilot string
	//go:embed TERMS_END_USER.md
	TermsEndUser string
	//go:embed TERMS_EXIT_NODE.md
	TermsExitNode string
	//go:embed TERMS_NODE_SHORT.md
	TermsNodeShort string
	//go:embed WARRANTY.md
	Warranty string
)
