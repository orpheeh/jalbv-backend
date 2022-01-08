package util

import (
	"crypto/sha256"
	"fmt"
)

func Hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func CompareHash(text, hash string) bool {
	return Hash(text) == hash
}
