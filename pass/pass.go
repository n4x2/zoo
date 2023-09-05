// Package pass provides password generation.
package pass

import (
	"bytes"
	"math/rand"
)

var (
	letter = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYX")
	number = []byte("0123456789")
	symbol = []byte("!@#$%^&*")
)

var easy = bytes.Join([][]byte{letter, number}, nil)

func Generate(l int, s bool) string {
	var p = make([]rune, l)

	if s {
		easy = append(easy, symbol...)
	}

	for i := range p {
		p[i] = rune(easy[rand.Intn(len(easy))])

	}

	return string(p)
}
