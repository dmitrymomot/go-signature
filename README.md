# go-signature

Just a small library to get a signed string with a payload.

## Usage

### Installation:
```golang
go get -u github.com/dmitrymomot/go-signature
```

### Set up signing key
```golang
import signature "github.com/dmitrymomot/go-signature"

...
signature.SigningKey = "your$ecret#key"
...
```

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