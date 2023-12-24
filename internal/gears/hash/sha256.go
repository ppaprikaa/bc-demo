package hash

import (
	"crypto/sha256"
)

func DoubleSHA256Hash(b []byte) []byte {
	f := sha256.Sum256(b)
	s := sha256.Sum256(f[:])
	return s[:]
}

func SHA256Hash(b []byte) []byte {
	f := sha256.Sum256(b)
	return f[:]
}
