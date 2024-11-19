package main

import (
	"fmt"
	"cryptographyServer/hashs"
)

func main() {
	var message string = "Hello World"

	hashedMsg, err := hashs.Hash(message, "sha512_224")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Hashed Message: ", hashedMsg)
}