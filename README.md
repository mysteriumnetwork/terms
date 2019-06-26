# Terms of use for Mysterium nodes

Provides terms and conditions for embedding into various platforms:
- `TERMS.md`
- `WARRANTY.md`

## Go

Regenerate after updating markdown sources:
```bash
go run mage.go -d terms-go generate
```

Usage:
```go
package main

import (
	"fmt"

	"github.com/mysteriumnetwork/terms/terms-go"
)

func main()  {
	fmt.Println(string(terms.WarrantyMdBytes))
	fmt.Println(string(terms.TermsMdBytes))
}
```
