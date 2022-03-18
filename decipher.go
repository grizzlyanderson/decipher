package main

//go:generate go run assets/generator.go

import (
	"flag"
	"fmt"
	"github.com/grizzlyanderson/decipher/calculators"
	"github.com/grizzlyanderson/decipher/cipher"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	log.SetLevel(log.DebugLevel)
	var inputfile, useCipher, key string
	var ignoreSpacess, textOnly bool
	flag.StringVar(&inputfile, "f", "", "source ciphertext file")
	flag.BoolVar(&ignoreSpacess, "s", true, "ingnore spaces in ciphertext")
	flag.BoolVar(&textOnly, "t", true, "ciphertext is ASCII only")
	flag.StringVar(&useCipher, "c", "", "encipher with named cipher, leave blank to attempt decipher [vignere|ceaser]")
	flag.StringVar(&key, "k", "", "key for enciphering")

	flag.Parse()

	if useCipher != "" {
		if key == "" {
			fmt.Println("A Key must be supplied when enciphering")
			return
		}
		doEncipher(useCipher, inputfile, key, ignoreSpacess)
		return
	}

	doDecypheryStuff(inputfile, ignoreSpacess)
}

func doEncipher(useCipher, inputfile, key string, ignoreSpaces bool) {
	// only one so whatever
	log.Info("using " + useCipher)
	plaintext, _ := ioutil.ReadFile(inputfile)

	switch useCipher {
	case "vignere":
		vc := cipher.NewVignere(key)
		// TODO should probably refactor encipher to take []byte
		prettyPrint(vc.Encipher(string(plaintext)))
	case "ceasar":
		if key == "all" {
			var s uint64 = 1
			for ; s <= 25; s++ {
				fmt.Printf("Shift by %v:", s)
				runRot(ignoreSpaces, plaintext, s)
			}
		} else {
			s, _ := strconv.ParseUint(key, 10, 8)
			runRot(ignoreSpaces, plaintext, s)
		}
	default:
		fmt.Printf("Unknown cipher type '%s'. Allowed types 'ceasar', 'vignere'\n", useCipher)
	}

}

func runRot(ignoreSpaces bool, plaintext []byte, s uint64) {
	if ignoreSpaces {
		cc := cipher.Rot(string(plaintext), uint8(s))
		prettyPrint(cc)
	} else {
		fmt.Println(cipher.ROTWithCase(plaintext, uint8(s)))
	}
}

func doDecypheryStuff(inputfile string, ignoreSpacess bool) {
	log.Debugf("Input file path: %s\n", inputfile)
	if "" == inputfile {
		flag.PrintDefaults()
		os.Exit(1)
	}
	// trivial read - not a good idea for larger files
	// also, blowing off err - very bad
	ciphertext, _ := ioutil.ReadFile(inputfile)
	log.Println(ciphertext)
	log.Println(string(ciphertext))
	charCounts, _ := calculators.CountByCharacters(ciphertext, ignoreSpacess)
	log.Println(charCounts)
	log.Println(len(charCounts))
	ic, _ := calculators.CalcIC(charCounts)
	log.Debugf("I.C. is %v.\n", ic)
	quadGram, err := calculators.LoadGrams(calculators.Eng, calculators.Quad)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(quadGram)
	log.Printf("number of quadgrams is %v\n", len(quadGram))
	log.Printf("Gram 'TION' count is %v\n", quadGram["TION"])
	setStates, err := calculators.GetNGramStats(calculators.Eng, calculators.Quad)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Gram 'TION' count is %v and probability is %v\n", setStates.NGramData["TION"].Count, setStates.NGramData["TION"].Probability)
	englishScore := calculators.Score([]byte("The quick brown fox jumped over the lazy dog Now is the time for all good englishment to come to the aid of their country To be or not to be, that is the question Weather tis nobler in the mind to suffer the slings and arrows of outragours misfortune"),
		setStates)
	exampleScore := calculators.Score([]byte("ATTACK THE EAST WALL OF THE CASTLE AT DAWN"), setStates)
	randoStats := calculators.Score([]byte(randStringBytesRmndr(256)), setStates)
	log.Printf("English language: %v\nRandom text: %v\nExample Score: %v\n", englishScore, randoStats, exampleScore)
	cipher.ShowPossiblePeriods(ciphertext, 15)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func prettyPrint(ciphertext string) {
	for i, c := range ciphertext {
		fmt.Print(string(c))
		if (i+1)%5 == 0 {
			fmt.Print(" ")
			if (i+1)%25 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
}
