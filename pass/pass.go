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

// invalidLengthError error returns if password length is invalid.
var invalidLengthError = errors.New("length must be between 1 and 50 characters")

// contains checks if any byte in subslice 'ss' is present in slice 's'.
func contains(s, ss []byte) bool {
	m := make(map[byte]bool, len(s))
	for _, c := range s {
		m[c] = true
	}

	for _, c := range ss {
		if m[c] {
			return true
		}
	}

	return false
}

// isConsecutiveType checks if two bytes have the same character type.
func isConsecutiveType(p, n byte) bool {
	var pr, nr = rune(p), rune(n)

	return (unicode.IsLower(pr) && unicode.IsLower(nr)) ||
		(unicode.IsUpper(pr) && unicode.IsUpper(nr)) ||
		(unicode.IsNumber(pr) && unicode.IsNumber(nr)) ||
		(unicode.IsSymbol(pr) && unicode.IsSymbol(nr))
}

// isRepeated checks if a byte 'b' is repeated within the given slice 's'.
func isRepeated(s []byte, b byte) bool {
	for _, v := range s {
		if v == b {
			return true
		}
	}

	return false
}

// randByte selects a random byte from a given slice of bytes.
func randByte(s []byte) byte {
	l := int64(len(s))

	i, err := rand.Int(rand.Reader, big.NewInt(l))
	if err != nil {
		panic(err)
	}

	return s[i.Int64()]
}

// Generate generates a random string. If no length is provided, it returns a default
// 20-character random string. If an invalid length is provided, it returns an error.
func Generate(n ...int) (string, error) {
	var l = defaultLength
	if len(n) > 0 {
		l = n[0]
	}

	if l < minLength || l > maxLength {
		return "", invalidLengthError
	}

	var p = make([]byte, l)
	var hasNumber, hasSymbol = false, false
	for i := range p {
		for {
			var c = randByte(chars)
			if i == 0 {
				p[i] = c
				break
			}

			if !isConsecutiveType(p[i-1], c) && !isRepeated(p[:i], c) {
				p[i] = c
				break
			}

			if !hasSymbol {
				if !contains(p[:i], symbols) {
					p[i] = randByte(symbols)
					hasSymbol = true
					break
				}
			}

			if !hasNumber {
				if !contains(p[:i], numbers) {
					p[i] = randByte(numbers)
					hasNumber = true
					break
				}
			}
		}
	}

	return string(p), nil
}
