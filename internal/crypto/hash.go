package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// GenerateHash take the route of the file and return the sha256 of the file.
// We use sha256 because it is a widely used and secure hashing algorithm.

func GenerateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// Return the hash as a hexadecimal string.
	return "0x" + hex.EncodeToString(hash.Sum(nil)), nil
}
