package calculators

func PeriodSplit(cipherchars []byte, period int) [][]byte {
	// create collections of bytes for each ordinal of the period
	cipherchars = purgeWhitespace(cipherchars)
	periodicChars := make([][]byte, period)
	for idx, char := range cipherchars {

		periodicChars[idx%period] = append(periodicChars[idx%period], char)
	}
	return periodicChars
}

func purgeWhitespace(cipherchars []byte) []byte {
	result := make([]byte, len(cipherchars))
	for _, c := range cipherchars {
		if !isSpace(c) {
			result = append(result, c)
		}
	}

	return result
}
