package main

import (
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
)

// Intended to be used with `go generate` to create ngram data files
func main() {
	// source or and information about ngamrs: http://practicalcryptography.com/cryptanalysis/text-characterisation/quadgrams/
	var fs http.FileSystem = http.Dir("data/english_quadgrams.txt")

	err := vfsgen.Generate(fs, vfsgen.Options{PackageName: "data", VariableName: "QuadGrams", Filename: "data/eng_quadgrams.go"})
	if err != nil {
		log.Fatalln(err)
	}
}
