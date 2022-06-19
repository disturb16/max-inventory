package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

const key string = "01234567890123456789012345678901"

func Encrypt(plaintext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func Decode(ciphertext string) (string, error) {
	bb, err := base64.RawStdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	decrypted, err := Decrypt(bb)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}
