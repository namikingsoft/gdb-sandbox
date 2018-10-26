package main

import (
	"fmt"
	"os"

	"app/src/lib"
)

func main() {
	plaintext := "Plain Text"
	if len(os.Args) > 1 {
		plaintext = os.Args[1]
	}

	ciphertext := lib.Encrypt(plaintext)
	plaintext = lib.Decrypt(ciphertext)

	fmt.Printf("%s => %s\n", plaintext, ciphertext)
}
