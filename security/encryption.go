package security

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/nacl/secretbox"
)

// Encrypt encrypts a report using a secret key
func Encrypt(report string, key [32]byte) (string, error) {
	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return "", err
	}

	encrypted := secretbox.Seal(nonce[:], []byte(report), &nonce, &key)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt decrypts a report using the same secret key
func Decrypt(encryptedData string, key [32]byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	var nonce [24]byte
	copy(nonce[:], data[:24])
	decrypted, ok := secretbox.Open(nil, data[24:], &nonce, &key)
	if !ok {
		return "", errors.New("decryption failed")
	}

	return string(decrypted), nil
}
