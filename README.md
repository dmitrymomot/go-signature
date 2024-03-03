# go-signature

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/dmitrymomot/go-signature)](https://github.com/dmitrymomot/go-signature)
[![Go Reference](https://pkg.go.dev/badge/github.com/dmitrymomot/go-signature.svg)](https://pkg.go.dev/github.com/dmitrymomot/go-signature)
[![License](https://img.shields.io/github/license/dmitrymomot/go-signature)](https://github.com/dmitrymomot/go-signature/blob/main/LICENSE)

[![Tests](https://github.com/dmitrymomot/go-signature/actions/workflows/tests.yml/badge.svg)](https://github.com/dmitrymomot/go-signature/actions/workflows/tests.yml)
[![CodeQL Analysis](https://github.com/dmitrymomot/go-signature/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dmitrymomot/go-signature/actions/workflows/codeql-analysis.yml)
[![GolangCI Lint](https://github.com/dmitrymomot/go-signature/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/dmitrymomot/go-signature/actions/workflows/golangci-lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmitrymomot/go-signature)](https://goreportcard.com/report/github.com/dmitrymomot/go-signature)

Just a small library to get a signed string with a payload.
It helps me to create confirmation tokens without using database.
> Don't use this library to protect sensitive data!

## Usage

### Installation:
```bash
go get -u github.com/dmitrymomot/go-signature
```

### Set up signing key
```golang
import signature "github.com/dmitrymomot/go-signature"

...
signature.SetSigningKey("secret-key")
...
```
Signing key will be set globally, so you don't need defining it each times

### Create signed string:
```golang
import signature "github.com/dmitrymomot/go-signature"

...
signedString, _ := signature.New("some data of any type")
log.Println(signedString)
...
```
Output:
```
eyJwIjoic29tZSBkYXRhIG9mIGFueSB0eXBlIn0.MzYwMzA0ZGVhNWRmMjdjOTM0ZjY1NzU3YWUwM2I0MDZmODRiMzRiMw
```

### Parse signed string:
```golang
data, err := signature.Parse("eyJwIjoic29tZSBkYXRhIG9mIGFueSB0eXBlIn0.MzYwMzA0ZGVhNWRmMjdjOTM0ZjY1NzU3YWUwM2I0MDZmODRiMzRiMw")
if err != nil {
    panic(err)
}
log.Println(data)
```
Output:
```
some data of any type
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
