package main

import (
	"os"
	"pixel_parser/src/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/test.ms")
	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}
}
