package main

//go:generate go run assets/generator.go

import (
	"flag"
	"fmt"
	"github.com/grizzlyanderson/decipher/calculators"
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

	charCounts, _ := calculators.CountByCharacters(ciphertext, ignoreSpacess)

	fmt.Println(charCounts)
	fmt.Println(len(charCounts))

	ic, _ := calculators.CalcIC(charCounts)
	fmt.Printf("I.C. is %v.\n", ic)
}
