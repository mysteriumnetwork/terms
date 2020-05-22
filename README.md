# Terms of use for BlockDEV AG (Mysterium Network) software

Provides terms and conditions for embedding into various platforms.

For local testing purposes you can generate after updating markdown sources (for all targets):
```bash
NEXT_VERSION=testver go run mage.go -v GenerateGo GenerateJs
```

Changes do not need to be committed. CI will generate changes automatically. 

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
