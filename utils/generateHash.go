package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateUserFolder(username, password, email string) string {
	data := username + password + email
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
