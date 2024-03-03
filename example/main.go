package main

import (
	"log"

	"github.com/dmitrymomot/go-signature"
)

type example struct {
	ID    uint64
	Email string
}

func main() {
	// set up signing key, it's highly important, don't forget!
	signature.SetSigningKey("secret-key")

	signedString, err := signature.New("some data of any type")
	if err != nil {
		panic(err)
	}
	log.Println("signed string", signedString)

	data, err := signature.Parse[string](signedString)
	if err != nil {
		panic(err)
	}
	log.Println(data)

	signedInt, err := signature.New(9834)
	if err != nil {
		panic(err)
	}
	log.Println("signed int", signedInt)

	siData, err := signature.Parse[int](signedInt)
	if err != nil {
		panic(err)
	}
	log.Println(siData)

	signedStruct, err := signature.New(example{ID: 123, Email: "test@m.dev"})
	if err != nil {
		panic(err)
	}
	log.Println("signed struct", signedStruct)

	ssData, err := signature.Parse[example](signedStruct)
	if err != nil {
		panic(err)
	}
	log.Println(ssData)
}
