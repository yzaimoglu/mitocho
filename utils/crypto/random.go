package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/yzaimoglu/mitocho/data/types"
)

func generateToken(tokenLength int) (string, error) {
	bytes := make([]byte, tokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateAccessToken() (types.AccessToken, error) {
	token, err := generateToken(64)
	return types.AccessToken(token), err
}
