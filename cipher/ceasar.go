package cipher

import (
	"github.com/decipher/calculators"
	"strings"
)

func Rot(plaintext string, shift uint8) string {
	plaintext = calculators.Normalize(plaintext)
	a, z := byte('A'), byte('Z')
	var result strings.Builder

	for _, c := range plaintext {
		b := byte(c)
		result.WriteByte((b-a+shift)%(z-a+1) + a)
	}
	return result.String()
}
