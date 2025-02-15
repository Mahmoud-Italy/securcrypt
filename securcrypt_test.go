package securcrypt_test

import (
	"testing"

	"github.com/mahmoud-italy/securcrypt"
)

func TestEncryptionDecryption(t *testing.T) {
	passphrase := "TestPassphrase"
	plaintext := "SecretData"

	encrypted, err := securcrypt.Encrypt(plaintext, passphrase)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := securcrypt.Decrypt(encrypted, passphrase)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if decrypted != plaintext {
		t.Fatalf("Expected %s but got %s", plaintext, decrypted)
	}
}
