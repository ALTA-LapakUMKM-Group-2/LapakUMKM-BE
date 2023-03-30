package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func EncryptText(plaintext string) string {
	key := []byte("F0rG3TpasswordKu")
	plaintextBytes := []byte(plaintext)
	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	nonce := make([]byte, 12)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return ""
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintextBytes, nil)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptText(ciphertext string) string {
	key := []byte("F0rG3TpasswordKu")
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return ""
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	// Create a new GCM (Galois Counter Mode) cipher
	nonce := make([]byte, 12)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return ""
	}

	// Decrypt the ciphertext using the GCM cipher
	plaintextBytes, err := aesgcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return ""
	}

	// Convert the decrypted plaintext byte slice to a string
	plaintext := string(plaintextBytes)

	// Return the decrypted plaintext string
	return plaintext
}
