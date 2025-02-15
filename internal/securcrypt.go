package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const (
	SaltSize  = 16
	KeySize   = 32
	NonceSize = 12
	Iter      = 4096
)

// DeriveKey generates a secure key from a given passphrase
func DeriveKey(passphrase, salt []byte) []byte {
	return pbkdf2.Key(passphrase, salt, Iter, KeySize, sha256.New)
}

// Encrypt encrypts plaintext using AES-GCM with a derived key
func Encrypt(plaintext string, passphrase string) (string, error) {
	salt := make([]byte, SaltSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	key := DeriveKey([]byte(passphrase), salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)
	finalData := append(salt, append(nonce, ciphertext...)...)
	return base64.StdEncoding.EncodeToString(finalData), nil
}

// Decrypt decrypts ciphertext using AES-GCM with a derived key
func Decrypt(encryptedText string, passphrase string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	if len(data) < SaltSize+NonceSize {
		return "", errors.New("invalid encrypted data")
	}

	salt := data[:SaltSize]
	nonce := data[SaltSize : SaltSize+NonceSize]
	ciphertext := data[SaltSize+NonceSize:]

	key := DeriveKey([]byte(passphrase), salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
