package main

import (
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
)

// Intended to be used with `go generate` to create ngram data files
func main() {
	// source or and information about ngamrs: http://practicalcryptography.com/cryptanalysis/text-characterisation/quadgrams/
	// specify individual files instead of a dir and each gofile will contain only that file, so no name requires when opening
	var fs http.FileSystem = http.Dir("assets")

	// TODO - if sticking with the dir path should probably no name the gofile for only 1 contained file/include other ngram files
	err := vfsgen.Generate(fs, vfsgen.Options{PackageName: "data", VariableName: "QuadGrams", Filename: "data/eng_quadgrams.go"})
	if err != nil {
		log.Fatalln(err)
	}
}
