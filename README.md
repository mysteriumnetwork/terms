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

[![npm version](https://badge.fury.io/js/%40mysteriumnetwork%2Fterms.svg)](https://badge.fury.io/js/%40mysteriumnetwork%2Fterms)

Usage:
```js
import {TermsEndUser, TermsExitNode} from '@mysteriumnetwork/terms'

console.log(TermsEndUser)
console.log(TermsExitNode)
```
