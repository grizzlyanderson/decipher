package cipher

import (
	"github.com/decipher/calculators"
	"strings"
)

const startPoint uint8 = 65 // English ASCII byte for letter A

// Vignere type for performing Vignere substitution en/deciphering
type Vignere struct {
	key string
	l   int
}

// NewVignere initializes a new struct with the enciphering key
func NewVignere(key string) *Vignere {
	return &Vignere{calculators.Normalize(key), len(key)}
}

// Encipher plain text to cipher text using Vignere substitution with the initialized key
// all letters will be converted to upper case, all spaces/non-letters will be removed, including numbers
func (v *Vignere) Encipher(plaintText string) string {
	pt := calculators.Normalize(plaintText)
	var ct strings.Builder

	for i, c := range pt {
		// get current letter of key
		kidx := i % v.l
		ki := a2i(v.key[kidx : kidx+1])
		pi := b2i(uint8(c))
		ci := (pi + ki) % 26
		ct.WriteByte(i2b(ci))
	}
	return ct.String()
}

// Decipher Vignere encrypted text using the initialized key
// result will be all uppercase with no spaces
func (v *Vignere) Decipher(cipherText string) string {
	ct := calculators.Normalize(cipherText)
	var pt strings.Builder

	for i, c := range ct {
		kidx := i % v.l
		ki := a2i(v.key[kidx : kidx+1])
		ci := b2i(uint8(c))
		// (cipher index + mod val - key index) mod val = plain index
		pi := (ci - ki + 26) % 26
		pt.WriteByte(i2b(pi))
	}
	return pt.String()
}

// convert alpha to integer (0-25); byte(uint8) to integer; or back; where integer is the - based ordinal of the letter in the alphabet
func a2i(c string) uint8 {
	return []byte(c)[0] - startPoint
}

func b2i(b uint8) uint8 {
	return b - startPoint
}

func i2a(i uint8) string {
	return string(i + startPoint)
}

func i2b(i uint8) uint8 {
	return i + startPoint
}
