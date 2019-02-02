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

func ROTWithCase(plaintext []byte, shift uint8) string {
	a, z, cA, cZ := byte('a'), byte('z'), byte('A'), byte('Z')
	var result strings.Builder
	for _, c := range plaintext {
		switch {
		case 'a' <= c && c <= 'z':
			result.WriteByte((c-a+shift)%(z-a+1) + a)
		case 'A' <= c && c <= 'Z':
			result.WriteByte((c-cA+shift)%(cZ-cA+1) + cA)
		default:
			result.WriteByte(c)
		}
	}
	return result.String()
}

