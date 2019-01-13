package calculators

// isSpace reports whether the byte is a space character.
// isSpace defines a space as being among the following bytes: ' ', '\t', '\n' and '\r'.
func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

// Get a count of each unique character in a []byte
// Whitespace characters are normally ignored but all other characters are counted
// For most uses the string should be normalized to uppercase, no punctuation or control characters
func CountByCharacters(cypherChars []byte, ignoreWitespace bool) (charCounts map[string]int, err error) {
	charCounts = make(map[string]int)
	for _, v := range cypherChars {
		if !ignoreWitespace || !isSpace(v) {
			charCounts[string(v)]++
		}
	}

	return charCounts, nil
}

// calculate the Index of Coincidence for map of counts by character from a ciphertext
// See http://practicalcryptography.com/cryptanalysis/text-characterisation/index-coincidence/ for information on I.C.
func CalcIC(counts map[string]int) (float64, error) {
	sum := 0
	totCount := 0
	for _, v := range counts {
		sum += v * (v - 1)
		totCount += v
	}
	ic := float64(sum) / float64(totCount*(totCount-1))

	return ic, nil
}
