package data

import _ "embed"

//go:embed ciphertext_vignere_CoreyDoctorow.txt
var CypherTextVignereBytes []byte

//go:embed ciphertext_vignere_CoreyDoctorow.txt
var CypherTextVignereString string

const VignereExampleKey = "CoreyDoctorow"
