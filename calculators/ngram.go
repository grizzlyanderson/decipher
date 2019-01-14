package calculators

import (
	"bufio"
	"bytes"
	"fmt"
	data2 "github.com/grizzlyanderson/decipher/data"
	"github.com/labstack/gommon/log"
	"io"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type lang int

// Enumeration of languages with ngram data
const (
	Eng lang = iota
)

func (l lang) toString() string {
	names := [...]string{
		"english",
	}
	return names[l]
}

type nGram int

// Enumeration of length of ngrams
const (
	Mono nGram = iota
	Bi
	Tri
	Quad
	Quint
)

func (g nGram) toString() string {
	names := [...]string{
		"mono",
		"bi",
		"tri",
		"quad",
		"quint",
	}
	return names[g]
}

// NGramStats provides a count and a probablity of that ngram occurring withing the complete set of ngrams
type NGramStats struct {
	Count       int
	Probability float64
}

// NGramCollection provides a map of stats for all the ngrams with the NGram as the key of the collection, as well as overall stats about the collection.
// including the GramLength (number of NGrams)
// ItemCOunt (sum of counts of all NGrams)
// Floor probability stat for the collection
type NGramCollection struct {
	NGramData  map[string]NGramStats
	GramLength int
	ItemCount  int64
	Floor      float64
}

// LoadGrams loads NGram data from an embedded resource for the specified language and length of gram (e.g. BI gram contains 2 letters 'AA' 'AB' etc)
// Resources are managed in assets_generate
func LoadGrams(language lang, gramLength nGram) (map[string]int, error) {
	result := make(map[string]int)
	if language != Eng {
		return nil, fmt.Errorf("Language %s not yet supported (no NGram files)", language.toString())
	}
	if gramLength != Quad {
		return nil, fmt.Errorf("Only Quad grams currently supported")
	}
	gramfile, err := data2.QuadGrams.Open(fmt.Sprintf("%s_%sgrams.txt", language.toString(), gramLength.toString()))
	if err != nil {
		return nil, err
	}
	defer gramfile.Close()

	// use bufio to read line by line
	// split the line on the space - characters before the space, count after
	reader := bufio.NewReader(gramfile)

	// line-level loop
	for {
		var (
			buffer   bytes.Buffer
			l        []byte // don't redeclare in loop
			isPrefix bool
			err      error // need to hang on to errout outside of the loop scope
		)

		// handle possible multi-part line (very very unlikely for this data)
		for {
			l, isPrefix, err = reader.ReadLine()
			if err != nil {
				break
			}

			buffer.Write(l)

			if !isPrefix {
				break
			}
		}
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		line := strings.Split(buffer.String(), " ")
		if len(line) != 2 {
			log.Errorf("Unexcepted number of line parts found. Expected 2, found %v in `%v`", len(line), buffer.String())
			break
		}
		r, err := strconv.ParseInt(line[1], 10, 32)
		if err != nil {
			log.Errorf("Unparsable number in line `%s`: %v", buffer.String(), err)
			err = nil
			break
		}
		result[line[0]] = int(r)
	}
	return result, nil
}

//GetNGramStats builds an NGramCollection for an NGram data set
func GetNGramStats(language lang, ngramType nGram) (NGramCollection, error) {
	gramMap, err := LoadGrams(language, ngramType)
	gramCollection := NGramCollection{NGramData: make(map[string]NGramStats)}
	if err != nil {
		return gramCollection, err
	}

	// whole collection info
	gramCollection.GramLength = int(ngramType) + 1
	//gramCollection.NGramData = make(map[string]NGramStats)
	for _, count := range gramMap {
		gramCollection.ItemCount += int64(count)
	}

	gramCollection.Floor = math.Log10(0.01 / float64(gramCollection.ItemCount))

	// probability for each gram in collection
	for k, count := range gramMap {
		//log.Printf("Gram: %s  Stat Count: %v   Total Item Count: %v", k, gramStat.Count, gramCollection.ItemCount)
		value := float64(count) / float64(gramCollection.ItemCount)
		prob := math.Log10(value)
		gramCollection.NGramData[k] = NGramStats{Count: count, Probability: prob}
	}

	return gramCollection, nil
}

// Score a sample text (clear or cipher) against an ngram set. More negative scores indicate more randomness,
// less negative scores indicate text that is much closer to the language of the ngram set
func Score(textToEvaluate []byte, ngrams NGramCollection) float64 {
	normalizedText := normalize(textToEvaluate)
	score := 0.0
	for i := 0; i < len(normalizedText)-ngrams.GramLength+1; i++ {
		if ng, ok := ngrams.NGramData[string(normalizedText[i:i+ngrams.GramLength])]; ok {
			score += ng.Probability
		} else {
			score += ngrams.Floor
		}
	}
	return score
}

func normalize(text []byte) []byte {
	result := make([]byte, 0)

	for _, b := range text {
		r := rune(b)
		if unicode.IsLetter(r) {
			result = append(result, byte(unicode.ToUpper(r)))
		}
	}

	return result
}
