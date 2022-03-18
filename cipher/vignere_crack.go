package cipher

import (
	"github.com/grizzlyanderson/decipher/calculators"
	log "github.com/sirupsen/logrus"
	"math"
)

func FindKeyLengths(cipherchars []byte, period int) {
	// TODO - identify likely targests and return lengths by looking for local maxes, ideally that occur a intervals of N
}

func ShowPossiblePeriods(cipherchars []byte, maxPeriod int) {
	for i := 2; i <= maxPeriod; i++ {
		log.Debugf("Period %v: %v\n", i, PeriodIC(cipherchars, i))
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
		v, err := calculators.CalcIC(c)
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
