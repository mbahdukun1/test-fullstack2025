package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1Hash(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
