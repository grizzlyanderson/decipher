package calculators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeriodSplit(t *testing.T) {
	text := []byte("A BCDEF GHIJ KLMN OP QRSTUVWXYZ")
	expected := map[int][]byte{
		0: []byte("ADGJMPSVY"),
		1: []byte("BEHKNQTWZ"),
		2: []byte("CFILORUX"),
	}

	result := PeriodSplit(text, 3)

	assert.Equal(t, expected, result)
}
