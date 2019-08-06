# Terms of use for Mysterium nodes

Provides terms and conditions for embedding into various platforms.

Regenerate after updating markdown sources (for all targets):
```bash
go run mage.go -v GenerateGo GenerateJs
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
	fmt.Println(string(terms.Warranty))
	fmt.Println(string(terms.TermsEndUser))
}
```

## Javascript

Usage:
```js
import {TermsEndUser, TermsExitNode} from 'terms-js'

console.log(TermsEndUser)
console.log(TermsExitNode)
```