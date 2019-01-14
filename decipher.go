package main

//go:generate go run assets/generator.go

import (
	"flag"
	"fmt"
	"github.com/grizzlyanderson/decipher/calculators"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"math/rand"
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

	quadGram, err := calculators.LoadGrams(calculators.Eng, calculators.Quad)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(quadGram)
	fmt.Printf("number of quadgrams is %v\n", len(quadGram))
	fmt.Printf("Gram 'TION' count is %v\n", quadGram["TION"])

	setStates, err := calculators.GetNGramStats(calculators.Eng, calculators.Quad)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gram 'TION' count is %v and probability is %v\n", setStates.NGramData["TION"].Count, setStates.NGramData["TION"].Probability)

	englishScore := calculators.Score([]byte("The quick brown fox jumped over the lazy dog Now is the time for all good englishment to come to the aid of their country To be or not to be, that is the question Weather tis nobler in the mind to suffer the slings and arrows of outragours misfortune"),
		setStates)

	exampleScore := calculators.Score([]byte("ATTACK THE EAST WALL OF THE CASTLE AT DAWN"), setStates)

	randoStats := calculators.Score([]byte(randStringBytesRmndr(256)), setStates)

	fmt.Printf("English language: %v\nRandom text: %v\nExample Sscore: %v", englishScore, randoStats, exampleScore)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
