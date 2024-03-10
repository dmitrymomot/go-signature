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
