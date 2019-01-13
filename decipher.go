package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var inputfile string
	var ignoreSpacess, textOnly bool
	flag.StringVar(&inputfile, "f", "", "source ciphertext file")
	flag.BoolVar(&ignoreSpacess, "s", true, "ingnore spaces in ciphertext")
	flag.BoolVar(&textOnly, "t", true, "ciphertext is ASCII only")

	flag.Parse()

	fmt.Println("Input file path: ", inputfile)
	if "" == inputfile {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// trivial read - not a good idea for larger files
	// also, blowing off err - very bad
	ciphertext, _ := ioutil.ReadFile(inputfile)

	fmt.Println(ciphertext)
	fmt.Println(string(ciphertext))

	charCounts, _ := countCharacters(ciphertext, ignoreSpacess)

	fmt.Println(charCounts)
	fmt.Println(len(charCounts))

	ic, _ := calcIC(charCounts)
	fmt.Println("I.C. is %v", ic)
}

// isSpace reports whether the byte is a space character.
// isSpace defines a space as being among the following bytes: ' ', '\t', '\n' and '\r'.
func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func countCharacters(cypherChars []byte, ignoreWitespace bool) (charCounts map[string]int, err error) {
	charCounts = make(map[string]int)
	for _, v := range cypherChars {
		if !ignoreWitespace || !isSpace(v) {
			charCounts[string(v)] += 1
		}
	}

	return charCounts, nil
}

func calcIC(counts map[string]int) (float64, error) {
	sum := 0
	totCount := 0
	for _, v := range counts {
		sum += v * (v - 1)
		totCount += v
	}
	ic := float64(sum) / float64(totCount*(totCount-1))

	return ic, nil
}
