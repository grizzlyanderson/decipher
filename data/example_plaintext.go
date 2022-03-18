package data

import _ "embed"

//go:embed plaintext.txt
var PlainTextBytes []byte

//go:embed plaintext.txt
var PlainTextString string
