package cipher

import (
	"fmt"
	"github.com/grizzlyanderson/decipher/calculators"
	log "github.com/sirupsen/logrus"
	"math"
	"strings"
)

func FindKeyLengths(cipherchars []byte, period int) {
	// TODO - identify likely targests and return lengths by looking for local maxes, ideally that occur a intervals of N
}

func GetPossiblePeriods(cipherchars []byte, maxPeriod int) (result map[int]float64) {
	// TODO - use dynamic channels & assemble back to map keyed on period length
	result = make(map[int]float64)
	for i := 1; i <= maxPeriod; i++ {
		ic := PeriodIC(cipherchars, i)

		log.Debugf("Period %v: %v\n", i, ic)
		result[i] = ic
	}
	return result
}

func ShowPossiblePeriods(ciperchars []byte, maxPeriod int) {
	periodICs := GetPossiblePeriods(ciperchars, maxPeriod)
	fmt.Printf("key: %-25s: val X 500\n", "IC Val")
	for i := 1; i <= len(periodICs); i++ {
		fmt.Printf("%-*v: %-25v: %s\n", 3, i, periodICs[i], strings.Repeat("x", int(periodICs[i]*500)))

	}

}

// PeriodIC splits the cipher text into a period of N, calculates the I.C. for each set and returns the avg I.C. for the entire period
func PeriodIC(cipherchars []byte, period int) float64 {
	periodicChars := calculators.PeriodSplit(cipherchars, period)

	var cumulatinveIC float64

	for _, periodChars := range periodicChars {
		c, err := calculators.CountByCharacters(periodChars, true)
		if err != nil {
			log.Error(err)
			return math.NaN()
		}
		v, err := calculators.CalcICForCharMap(c)
		if err != nil {
			log.Error(err)
			return math.NaN()
		}
		cumulatinveIC += v
		if cumulatinveIC == math.NaN() {
			log.Errorf("Oh shit:: periodChars: %v\n  v: %v\n  c: %v\n  err: %v\n", periodChars, v, c, err)
		}
	}
	// TODO - something is wrong - the values are coming out way too high, but not figuring it out tonight
	log.Debugf("cumulativeIC: %v      period: %v", cumulatinveIC, period)
	periodicIC := cumulatinveIC / float64(period)

	return periodicIC
}

func PeriodIC2(ciphertext []byte, period int) (float64, error) {

	// if period == 1, return plain IC
	if period == 1 {
		return calculators.CalcICForCyphertext(ciphertext)
	}

	// for periods 2 -> N, get N slices containing ever Nth character starting from 0th, to N-1th
	// eg period 3 0,3,6,9...; 1,4,7,10,...; 2,5,8,11...
	periodicCypherChars := calculators.PeriodSplit(ciphertext, period)
	sumIC := 0.0
	for _, cypherChars := range periodicCypherChars {
		ic, err := calculators.CalcICForCyphertext(cypherChars)
		if err != nil {
			return math.NaN(), err
		}
		log.Info(ic)
		sumIC += ic
	}

	return sumIC / float64(period), nil

}
