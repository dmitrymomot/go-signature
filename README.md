# go-signature

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/dmitrymomot/go-signature/v2)](https://github.com/dmitrymomot/go-signature/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/dmitrymomot/go-signature/v2.svg)](https://pkg.go.dev/github.com/dmitrymomot/go-signature/v2)
[![License](https://img.shields.io/github/license/dmitrymomot/go-signature/v2)](https://github.com/dmitrymomot/go-signature/v2/blob/main/LICENSE)

[![Tests](https://github.com/dmitrymomot/go-signature/v2/actions/workflows/tests.yml/badge.svg)](https://github.com/dmitrymomot/go-signature/v2/actions/workflows/tests.yml)
[![CodeQL Analysis](https://github.com/dmitrymomot/go-signature/v2/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dmitrymomot/go-signature/v2/actions/workflows/codeql-analysis.yml)
[![GolangCI Lint](https://github.com/dmitrymomot/go-signature/v2/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/dmitrymomot/go-signature/v2/actions/workflows/golangci-lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmitrymomot/go-signature/v2)](https://goreportcard.com/report/github.com/dmitrymomot/go-signature/v2)

Just a small library to get a signed string with a payload.
It helps me to create confirmation tokens without using database.
> Don't use this library to protect sensitive data!

## Features

- **Data Integrity**: Ensures that the data remains unchanged and secure during transit.
- **Simplified Token Structure**: Generates tokens without the overhead of JWT headers, focusing solely on payload and signature.
- **Flexibility and Ease of Use**: Provides a straightforward API to work with, requiring minimal setup to sign and verify data.

## Usage

### Installation:

```bash
go get -u github.com/dmitrymomot/go-signature/v2
```

### Example:

Use the Signer  to sign some predefined data type and parse the token back to the original data.

```golang
package main

import (
	"fmt"

	"github.com/dmitrymomot/go-signature/v2"
)

// Define a struct to sign. Or use any other data type you want.
type example struct {
	ID    uint64
	Email string
}

func main() {
	// Create a new signer
	s := signature.NewSigner[example]([]byte("signing-key"))

	// Sign and parse a token
	token, err := s.Sign(example{ID: 123, Email: "test123"})
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	// Parse a token and print the data
	data, err := s.Parse(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
```

Output:

```shell
eyJwIjp7IklEIjoxMjMsIkVtYWlsIjoidGVzdDEyMyJ9fQ.Mjg1MDVkOTNjNTdkNDhjMjk2NWQxOWZhNGY3ZDU2ZjQ3NWFlNWUxYw

{123 test123}
```

You can find this example in the [example/main.go](example/main.go) file.

#### Using of functions directly

You can use the `NewToken` and `ParseToken` functions directly without creating a new signer.

```golang
package main

import (
    "fmt"

    "github.com/dmitrymomot/go-signature/v2"
)

func main() {
    signingKey := []byte("signing-key")
    someData := "some data"

    token, err := signature.NewToken(signingKey, someData, 0, signature.CalculateHmac)
    if err != nil {
        panic(err)
    }
    fmt.Println(token)

    // Parse a token and print the data. You need to know the type of the data to parse it.
    data, err := signature.ParseToken[string](signingKey, token, signature.CalculateHmac, signature.ValidateHmac)
    if err != nil {
        panic(err)
    }
    fmt.Println(data)
}
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
