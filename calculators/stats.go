// Package calculators contains an assortment of calculation functions to manipulate []byte of
// ciphertext characters for use in different types of anlysis or to actually perform various
// statistical analysis.
package calculators

import "math"

// isSpace reports whether the byte is a space character.
// isSpace defines a space as being among the following bytes: ' ', '\t', '\n' and '\r'.
func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

// CountByCharacters gets a count of each unique character in a []byte
// Whitespace characters are normally ignored but all other characters are counted
// For most uses the string should be normalized to uppercase, no punctuation or control characters
func CountByCharacters(cypherChars []byte, ignoreWitespace bool) (charCounts map[byte]int, err error) {
	charCounts = make(map[byte]int)
	for _, v := range cypherChars {
		if !ignoreWitespace || !isSpace(v) {
			charCounts[v]++
		}
	}

	return charCounts, nil
}

// CalcICForCharMap returns Index of Coincidence for map of counts by character from a ciphertext
// See http://practicalcryptography.com/cryptanalysis/text-characterisation/index-coincidence/ for information on I.C.
func CalcICForCharMap(counts map[byte]int) (float64, error) {
	sum := 0
	totCount := 0

	// just in case characters other than letters are included, ignore them
	// this is still assuming that the text is all upper case...
	for i := byte('A'); i <= byte('Z'); i++ {
		if v, ok := counts[i]; ok {
			sum += v * (v - 1)
			totCount += v
		}
	}

	ic := float64(sum) / float64(totCount*(totCount-1))

	return ic, nil
}

func CalcICForCiphertext(ciphertext []byte) (float64, error) {
	if charMap, err := CountByCharacters(ciphertext, true); err != nil {
		return math.NaN(), err
	} else {
		return CalcICForCharMap(charMap)
	}
}
