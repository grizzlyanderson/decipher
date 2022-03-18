package cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVignere_EncipherTrivialKey(t *testing.T) {
	key := "A"
	plainText := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expectedText := "BCDEFGHIJKLMNOPQRSTUVWXYZA"

	vc := NewVignere(key)
	cipherText := vc.Encipher(plainText)

	//assert.Equalf(t, expectedText, cipherText, "expected '%s', got '%s'", expectedText, cipherText)
	assert.Equal(t, expectedText, cipherText)
}

func TestVignere_EncipherLongishKey(t *testing.T) {
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	plainText := "ONE WHITE DUCK"
	expectedText := "PPHAMOAMMENW"

	vc := NewVignere(key)
	cipherText := vc.Encipher(plainText)

	assert.Equal(t, expectedText, cipherText)
}

func TestVignere_Decipher(t *testing.T) {
	key := "A"
	expectedText := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cipherText := "BCDEFGHIJKLMNOPQRSTUVWXYZA"

	vc := NewVignere(key)
	plainText := vc.Decipher(cipherText)

	assert.Equal(t, expectedText, plainText)
}

func TestVignere_DecipherLongishKey(t *testing.T) {
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expectedText := "ONEWHITEDUCK"
	cipherText := "PPHAMOAMMENW"

	vc := NewVignere(key)
	plainText := vc.Decipher(cipherText)

	assert.Equal(t, expectedText, plainText)
}
