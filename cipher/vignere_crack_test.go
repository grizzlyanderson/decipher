package cipher

import (
	"github.com/grizzlyanderson/decipher/calculators"
	"github.com/grizzlyanderson/decipher/data"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeriodICPeriodOf1(t *testing.T) {
	cyphertext := []byte(calculators.Normalize(data.PlainTextString))
	expectedIC, _ := calculators.CalcICForCiphertext(cyphertext)

	calculatedIC, err := PeriodIC(cyphertext, 1)
	assert.Nil(t, err)
	assert.Equal(t, expectedIC, calculatedIC)
}

func TestPeriodICPeriodOf3PlainEnglish(t *testing.T) {
	period := 3
	cyphertext := []byte(calculators.Normalize(data.PlainTextString))
	wholeIC, _ := calculators.CalcICForCiphertext(cyphertext)
	log.Infof("WholeIC: %v", wholeIC)

	calculatedIC, err := PeriodIC(cyphertext, period)
	assert.Nil(t, err)
	assert.Truef(t,
		0.98*wholeIC < calculatedIC && calculatedIC < 1.02*wholeIC,
		"Calculated IC expected to be w/in +/- 2% of %n", wholeIC)
}

func TestGetPossiblePeriods(t *testing.T) {
	cyphertext := []byte(calculators.Normalize(data.CypherTextVignereString))
	expectedKeyLength := len(data.VignereExampleKey)
	periods := int(float64(expectedKeyLength) * 1.5)
	periodICs := GetPossiblePeriods(cyphertext, periods)

	for period, ic := range periodICs {
		if period != expectedKeyLength {
			assert.Less(t, ic, periodICs[expectedKeyLength])
		}
	}
}

func TestShowPossiblePeriods(t *testing.T) {
	t.SkipNow() // not really a test, just a short cut to seeing the visual output
	log.SetLevel(log.InfoLevel)
	cyphertext := []byte(calculators.Normalize(data.CypherTextVignereString))
	expectedKeyLength := len(data.VignereExampleKey)
	periods := int(float64(expectedKeyLength) * 1.5)
	ShowPossiblePeriods(cyphertext, periods)
}

func TestGuessVignereKeyLength(t *testing.T) {
	periodicICs := map[int]float64{
		1: 0.001,
		2: 0.0122,
		3: 0.0132,
		4: 0.047,
		5: 0.025,
		6: 0.0188,
	}

	keylength := GuessVignereKeyLength(periodicICs)
	assert.Equal(t, 4, keylength)
}
