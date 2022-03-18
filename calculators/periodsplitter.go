package calculators

func PeriodSplit(cipherchars []byte, period int) map[int][]byte {
	// create collections of bytes for each ordinal of the period
	cleaned := purgeWhitespace(cipherchars)
	periodicChars := make(map[int][]byte)
	for idx, char := range cleaned {

		periodicChars[idx%period] = append(periodicChars[idx%period], char)
	}
	return periodicChars
}
