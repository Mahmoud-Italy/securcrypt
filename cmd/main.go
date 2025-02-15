package main

import (
	"fmt"
	"log"

	"github.com/mahmoud-italy/securcrypt"
)

func main() {
	passphrase := "SuperSecurePassword123!"
	plaintext := "Hello, this is a secret message."

	// Encrypt the message
	encryptedText, err := securcrypt.Encrypt(plaintext, passphrase)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}

	fmt.Println("Encrypted:", encryptedText)

	// Decrypt the message
	decryptedText, err := securcrypt.Decrypt(encryptedText, passphrase)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	fmt.Println("Decrypted:", decryptedText)
}
