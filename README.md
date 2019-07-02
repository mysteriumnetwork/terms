# Terms of use for Mysterium nodes

Provides terms and conditions for embedding into various platforms:
- `TERMS.md`
- `WARRANTY.md`

Regenerate after updating markdown sources (for all targets):
```bash
go run mage.go generate
```

## Go

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

## Javascript

Usage:
```js
import {TermsMd, WarrantyMd} from 'terms-js'

console.log(TermsMd)
console.log(WarrantyMd)
```