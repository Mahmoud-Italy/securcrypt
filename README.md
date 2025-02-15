# SecureCrypt - AES-GCM Encryption in Go

SecureCrypt is a Go package that provides robust encryption and decryption using AES-GCM with a secure key derivation function (PBKDF2). It ensures high security, authenticity, and integrity of your encrypted data.

## Features

- AES-GCM Encryption (Authenticated Encryption)

- PBKDF2 Key Derivation (Strengthens password-based encryption)

- Random Salt & Nonce Generation (Ensures uniqueness)

- Base64 Encoding (Safe for storage & transmission)

- Secure and Efficient Implementation

## Installation

```bash
composer go get -u github.com/mahmoud-italy/securecrypt
```

## Usage
### Import the package

```bash
import (
    "fmt"
    "log"
    "github.com/mahmoud-italy/securecrypt"
)
```

### Encrypting a Message
```bash
passphrase := "SuperSecurePassword123!"
plaintext := "Hello, this is a secret message."

encryptedText, err := securecrypt.Encrypt(plaintext, passphrase)
if err != nil {
    log.Fatalf("Encryption failed: %v", err)
}
fmt.Println("Encrypted:", encryptedText)
```

### Decrypting a Message

```bash
decryptedText, err := securecrypt.Decrypt(encryptedText, passphrase)
if err != nil {
    log.Fatalf("Decryption failed: %v", err)
}
fmt.Println("Decrypted:", decryptedText)
```

## Running Tests
We have included test cases to ensure the correctness of the encryption and decryption logic.
### Run tests using:
```bash
go test ./...
```

### Example Test Case (Inside securecrypt_test.go)
```bash
package securecrypt

import (
    "testing"
)

func TestEncryptionDecryption(t *testing.T) {
    passphrase := "TestPassphrase"
    plaintext := "SecretData"

    encrypted, err := Encrypt(plaintext, passphrase)
    if err != nil {
        t.Fatalf("Encryption failed: %v", err)
    }

    decrypted, err := Decrypt(encrypted, passphrase)
    if err != nil {
        t.Fatalf("Decryption failed: %v", err)
    }

    if decrypted != plaintext {
        t.Fatalf("Expected %s but got %s", plaintext, decrypted)
    }
}
```

## License
The MIT License (MIT). Please see [License File](LICENSE.md) for more information.

## Contributing
Feel free to contribute by submitting issues or pull requests.