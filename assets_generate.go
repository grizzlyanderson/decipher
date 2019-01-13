package main

import (
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
)

func main() {
	var fs http.FileSystem = http.Dir("data/english_quadgrams.txt")

	err := vfsgen.Generate(fs, vfsgen.Options{PackageName: "data", VariableName: "QuadGrams", Filename: "data/eng_quadgrams.go"})
	if err != nil {
		log.Fatalln(err)
	}
}
