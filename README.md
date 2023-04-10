# go-signature

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
signature.SetSigningKey("secret%key")
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

---

Licensed under [Apache License 2.0](https://github.com/dmitrymomot/go-signature/blob/master/LICENSE)
