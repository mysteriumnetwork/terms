package generate

import (
	"github.com/magefile/mage/mg"
)

// Generate generates packages for all language targets
func Generate() error {
	mg.SerialDeps(
		GenerateGo,
		GenerateJs,
	)
	return nil
}
