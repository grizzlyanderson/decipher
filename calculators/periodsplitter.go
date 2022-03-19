package calculators

// PeriodSplit breaks a []byte] (usually ciphertext characters) into as many []byte as the value of period
// each []byte contains every Nth character. The map uses step as the key, so for a period of 3
// there will be map[1] which contains character 1,4,7,10..., map[2] 2,5,8,11... and map[0] 3,6,9,12,...
func PeriodSplit(cipherchars []byte, period int) map[int][]byte {
	// create collections of bytes for each ordinal of the period
	cleaned := purgeWhitespace(cipherchars)
	periodicChars := make(map[int][]byte)
	for idx, char := range cleaned {

		periodicChars[idx%period] = append(periodicChars[idx%period], char)
	}
	return periodicChars
}
