package data

import _ "embed"

//go:embed plaintext.txt
var CypherTextVignereBytes []byte

//go:embed plaintext.txt
var CypherTextVignereString string

const VignereExampleKey = "CoreyDoctrow"
