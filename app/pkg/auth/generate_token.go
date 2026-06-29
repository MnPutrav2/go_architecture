package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"
)

func GenerateRefreshToken() (string, string, time.Time, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", "", time.Time{}, err
	}

	exp := time.Now().Add(10)

	token := base64.RawURLEncoding.EncodeToString(b)

	return token, HashToken(token), exp, nil
}

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
