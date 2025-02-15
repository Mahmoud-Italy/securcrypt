package test

import (
	"testing"

	"github.com/mahmoud-italy/securcrypt/internal"
)

func TestEncryptionDecryption(t *testing.T) {
	passphrase := "TestPassphrase"
	plaintext := "SecretData"

	encrypted, err := internal.Encrypt(plaintext, passphrase)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := internal.Decrypt(encrypted, passphrase)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if decrypted != plaintext {
		t.Fatalf("Expected %s but got %s", plaintext, decrypted)
	}
}
