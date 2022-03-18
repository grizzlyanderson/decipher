package calculators

import (
	"github.com/grizzlyanderson/decipher/data"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCountByCharacters(t *testing.T) {
	chars := []byte("ABBCCCDDDDEEEEEFFFFFFGGGGGGG")
	expectedResult := map[byte]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'D': 4,
		'E': 5,
		'F': 6,
		'G': 7,
	}

	result, e := CountByCharacters(chars, true)
	assert.Nil(t, e)
	assert.Equal(t, expectedResult, result)
}

func TestCountByCharacters_IgnoreWhiteSpace(t *testing.T) {
	chars := []byte(`A
BB
CCC DDDD
EEEEE   FFFFFF
GGGGGGG`)

	expectedResult := map[byte]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'D': 4,
		'E': 5,
		'F': 6,
		'G': 7,
	}

	result, e := CountByCharacters(chars, true)
	assert.Nil(t, e)
	assert.Equal(t, expectedResult, result)
}
func TestCountByCharacters_CountWhiteSpace(t *testing.T) {
	chars := []byte(`A
BB
CCC DDDD
EEEEE FFFFFF
GGGGGGG`)

	expectedResult := map[byte]int{
		'A':  1,
		'B':  2,
		'C':  3,
		'D':  4,
		'E':  5,
		'F':  6,
		'G':  7,
		'\n': 4,
		' ':  2,
	}

	result, e := CountByCharacters(chars, false)
	assert.Nil(t, e)
	assert.Truef(t, reflect.DeepEqual(expectedResult, result), "Not Equal.\nexpected: %v \nactual  : %v", expectedResult, result)
}

func TestCalcIC(t *testing.T) {
	expectedHigh := 0.0684
	expectedLow := 0.0682
	charCount, _ := CountByCharacters([]byte(Normalize(data.PlainTextString)), true)
	ic, e := CalcIC(charCount)

	assert.Nil(t, e)
	// getting slightly different values calculating ic on practicalcryptography, other sources so expect within a narrow range.
	assert.Truef(t, expectedLow <= ic && ic <= expectedHigh, "Out of range: Expected %v < %v < %v", expectedLow, ic, expectedHigh)
}
