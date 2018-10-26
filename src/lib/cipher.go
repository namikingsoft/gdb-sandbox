package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func newCipher() cipher.Block {
	keytext := os.Getenv("CIPHER_KEY_PHRASE")
	c, err := aes.NewCipher([]byte(keytext))
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(-1)
	}
	return c
}

func Encrypt(plaintext string) string {
	plaintextBytes := []byte(plaintext)
	c := newCipher()
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertextBytes := make([]byte, len(plaintextBytes))
	cfb.XORKeyStream(ciphertextBytes, plaintextBytes)
	return fmt.Sprintf("%x", ciphertextBytes)
}

func Decrypt(ciphertext string) string {
	ciphertextBytes, _ := hex.DecodeString(ciphertext)
	c := newCipher()
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextBytes := make([]byte, len(ciphertextBytes))
	cfbdec.XORKeyStream(plaintextBytes, ciphertextBytes)
	return string(plaintextBytes)
}
