package hash

import (
	"crypto/sha256"
	"fmt"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

func GeneratePasswordHash(password string) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(salt))), nil
}
