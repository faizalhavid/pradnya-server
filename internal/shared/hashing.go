package shared

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashData(data string) (string, error) {
	hash := sha256.New()

	_, err := hash.Write([]byte(data))
	if err != nil {
		return "", err
	}
	hashedData := hash.Sum(nil)
	return hex.EncodeToString(hashedData), nil
}
