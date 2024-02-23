package crypto

import (
	"crypto/rand"
	"encoding/hex"
)

func generateToken(tokenLength int) (string, error) {
	bytes := make([]byte, tokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateAccessToken() (string, error) {
	return generateToken(64)
}
