// Package pass provides password generation.
package pass

import (
	"crypto/rand"
	"errors"
	"math/big"
	"unicode"

	"github.com/n4x2/zoo/is"
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
)

// invalidLengthError error returns if password length is invalid.
var invalidLengthError = errors.New("length must be between 1 and 50 characters")

// isConsecutiveType checks if two bytes have the same character type.
func isConsecutiveType(p, n byte) bool {
	var pr, nr = rune(p), rune(n)

	return (unicode.IsLower(pr) && unicode.IsLower(nr)) ||
		(unicode.IsUpper(pr) && unicode.IsUpper(nr)) ||
		(unicode.IsNumber(pr) && unicode.IsNumber(nr)) ||
		(unicode.IsSymbol(pr) && unicode.IsSymbol(nr))
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

// Generate generates a random string.
//
// It returns a default 20-character random string containing letters, symbols and numbers.
// Symbols and numbers is optional, set 'n' to false to exclude numbers or 's' to false to
// exclude symbols.
//
// It will returns an error if provided length 'l' less than 1 or greater than 50.
func Generate(n, s bool, l ...int) (string, error) {
	var length = defaultLength
	if len(l) > 0 {
		length = l[0]
	}

	if length < minLength || length > maxLength {
		return "", invalidLengthError
	}

	var chars = letters
	var hasNumber, hasSymbol = true, true
	if n {
		chars = append(chars, numbers...)
		hasNumber = false
	}

	if s {
		chars = append(chars, symbols...)
		hasSymbol = false
	}

	var p = make([]byte, length)
	for i := range p {
		for {
			var c = randByte(chars)
			if i == 0 {
				p[i] = c
				break
			}

			if !isConsecutiveType(p[i-1], c) && !is.Contain(p[:i], c) {
				p[i] = c
				break
			}

			if !hasSymbol {
				if !is.ContainOneOf(p[:i], symbols) {
					p[i] = randByte(symbols)
					hasSymbol = true
					break
				}
			}

			if !hasNumber {
				if !is.ContainOneOf(p[:i], numbers) {
					p[i] = randByte(numbers)
					hasNumber = true
					break
				}
			}
		}
	}

	return string(p), nil
}
