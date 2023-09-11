// Package pass provides password generation.
package pass

import (
	"bytes"
	"crypto/rand"
	"errors"
	"math/big"
	"unicode"
)

const (
	defaultLength = 20
	minLength     = 1
	maxLength     = 50
)

var (
	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers = []byte("0123456789")
	symbols = []byte("!@#$%^&*")
	chars   = bytes.Join([][]byte{letters, numbers, symbols}, nil)
)

// invalidLengthError is the error returned when the password length is invalid.
var invalidLengthError = errors.New("password length invalid, must be between 1 and 50 characters")

// randByte selects a random byte from a given slice of bytes.
func randByte(v []byte) byte {
	l := int64(len(v))

	i, err := rand.Int(rand.Reader, big.NewInt(l))
	if err != nil {
		panic(err)
	}

	return v[i.Int64()]
}

// isConsecutiveType checks if two bytes have the same character type.
func isConsecutiveType(p, n byte) bool {
	if (unicode.IsLower(rune(p)) && unicode.IsLower(rune(n))) ||
		(unicode.IsUpper(rune(p)) && unicode.IsUpper(rune(n))) ||
		(unicode.IsNumber(rune(p)) && unicode.IsNumber(rune(n))) ||
		(unicode.IsSymbol(rune(p)) && unicode.IsSymbol(rune(n))) {
		return true
	}
	return false
}

// contains checks if any byte in slice 'v' is present in slice 'p'.
func contains(p, v []byte) bool {
	for _, c := range v {
		for _, val := range p {
			if val == c {
				return true
			}
		}
	}

	return false
}

// isRepeated checks if a byte 'n' is repeated within the given slice 'p'.
func isRepeated(p []byte, n byte) bool {
	for _, v := range p {
		if v == n {
			return true
		}
	}

	return false
}

// Generate generates a random string of a specific length. If no length is provided,
// it returns a default 20-character random string. If an invalid length is provided,
// it returns an error.
func Generate(n ...int) (string, error) {
	var l int
	if len(n) > 0 {
		l = n[0]
	} else {
		l = defaultLength
	}

	if l < minLength || l > maxLength {
		return "", invalidLengthError
	}

	var p = make([]byte, l)
	for i := range p {
		for {
			c := randByte(chars)
			if i == 0 {
				p[i] = c
				break
			}

			if !contains(p[:i], symbols) {
				c = randByte(symbols)
			}

			if !contains(p[:i], numbers) {
				c = randByte(numbers)
			}

			if i != 0 && !isConsecutiveType(p[i-1], c) && !isRepeated(p[:i], c) {
				p[i] = c
				break
			}
		}
	}

	return string(p), nil
}
