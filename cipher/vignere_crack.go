package cipher

import (
	"fmt"
	"github.com/grizzlyanderson/decipher/calculators"
	log "github.com/sirupsen/logrus"
	"math"
	"strings"
)

func GuessVignereKeyLength(periodICs map[int]float64) (keylength int) {
	maxIC := 0.0
	keylength = int(math.NaN())
	for k, ic := range periodICs {
		if ic > maxIC {
			maxIC = ic
			keylength = k
		}
	}
	return keylength
}

// GetPossiblePeriods is primarily used to guess the length of key for a vignere cipher.
// It takes a maxPeriod (key length) and returns a map of 2 to maxPeriod and returns a collection
// of average IC values keyed by possible key length. Generally the key with the largest value is the
// length of the key used in the vignere cipher
func GetPossiblePeriods(cipherchars []byte, maxPeriod int) (result map[int]float64) {
	// TODO - use dynamic channels & assemble back to map keyed on period length
	result = make(map[int]float64)
	for i := 1; i <= maxPeriod; i++ {
		ic, err := PeriodIC(cipherchars, i)

		if err != nil {
			log.Error(err)
			result[i] = 0.0
			continue
		}

		log.Debugf("Period %v: %v\n", i, ic)
		result[i] = ic
	}
	return result
}

// ShowPossiblePeriods is a display function that turns a map of keyLength:IC values
// into a visual display, a horizontal bar-chart of the IC Values by key length
func ShowPossiblePeriods(ciperchars []byte, maxPeriod int) {
	periodICs := GetPossiblePeriods(ciperchars, maxPeriod)
	fmt.Printf("key: %-25s: val X 500\n", "IC Val")
	for i := 1; i <= len(periodICs); i++ {
		fmt.Printf("%-*v: %-25v: %s\n", 3, i, periodICs[i], strings.Repeat("x", int(periodICs[i]*500)))
	}

}

// PeriodIC splits the cipher text into a period of N, calculates the I.C. for each set and returns the avg I.C. for the entire period
func PeriodIC(ciphertext []byte, period int) (float64, error) {

	// if period == 1, return plain IC
	if period == 1 {
		return calculators.CalcICForCiphertext(ciphertext)
	}

	// for periods 2 -> N, get N slices containing ever Nth character starting from 0th, to N-1th
	// eg period 3 0,3,6,9...; 1,4,7,10,...; 2,5,8,11...
	periodicCypherChars := calculators.PeriodSplit(ciphertext, period)
	sumIC := 0.0
	for _, cypherChars := range periodicCypherChars {
		ic, err := calculators.CalcICForCiphertext(cypherChars)
		if err != nil {
			return math.NaN(), err
		}
		log.Debug(ic)
		sumIC += ic
	}

	return sumIC / float64(period), nil
}
