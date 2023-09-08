// Package pass provides password generation.
package pass

import (
	"crypto/rand"
	"errors"
	"math/big"
	"slices"
)

var (
	chars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*")
)

// errInvalidLength error return if length less than or equal zero.
var errInvalidLength = errors.New("password length is invalid: ideally should be at least 12 characters long")

// Generate generate random string in specific length.
//
// TODO
// - Avoid repeated characters
// - Avoid consecutive characters
// - Make sure at least one symbol in random string
func Generate(l int) (string, error) {
	if l <= 0 {
		return "", errInvalidLength
	}

	var p = make([]byte, l)

	for i := range p {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		p[i] = chars[n.Int64()]

		chars = slices.Delete(chars, int(n.Int64()), int(n.Int64()+1))
	}

	return string(p), nil
}
