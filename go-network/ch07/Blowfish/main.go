package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/blowfish"
)

// go get golang.org/x/crypto/blowfish
func main() {
	key := []byte("my key")
	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	src := []byte("Show me the money, plz!\n\n\n")
	var enc [512]byte

	cipher.Encrypt(enc[0:], src)

	var descrypt [32]byte
	cipher.Decrypt(descrypt[0:], enc[0:])
	result := bytes.NewBuffer(nil)
	result.Write(descrypt[0:32])
	fmt.Println(string(result.Bytes())) // Show me
}
