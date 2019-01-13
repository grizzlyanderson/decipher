package calculators

import (
	"bufio"
	"bytes"
	"fmt"
	data2 "github.com/grizzlyanderson/decipher/data"
	"github.com/labstack/gommon/log"
	"io"
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

// need to use code from : https://web.archive.org/web/20180424121229/http://practicalcryptography.com:80/cryptanalysis/text-characterisation/quadgrams/#a-python-implementation

// should load wuad grams, but might encode all n grams for english

// also might build n-gram loader that can be fed files from project guttenberg

// and consider https://github.com/go-bindata/go-bindata to create & embed resource files

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

func normalize(ciphertext []byte) []byte {
	result := make([]byte, 0)

	for _, b := range ciphertext {
		r := rune(b)
		if unicode.IsLetter(r) {
			result = append(result, byte(unicode.ToUpper(r)))
		}
	}

	return result
}
