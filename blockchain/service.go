package blockchain

import (
	"crypto/sha256"
	"fmt"
)

// GenerateHash creates a SHA-256 hash of the report
func GenerateHash(report string) string {
	hash := sha256.Sum256([]byte(report))
	return fmt.Sprintf("%x", hash)
}
