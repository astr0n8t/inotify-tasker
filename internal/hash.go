package internal

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// Either returns f if m is filename
// or hashes the file that f represents and returns the sha256 hash as a string
//
// If there is an issue simply returns the hash as an empty string
func hash(f string, m string) (string, error) {
	if m == "hash" {
		file, err := os.ReadFile(f)
		if err != nil {
			return "", fmt.Errorf("issue while hashing %v : %v", f, err)
		}
		return fmt.Sprintf("%x", sha256.Sum256(file)), nil
	}
	return f, nil
}
