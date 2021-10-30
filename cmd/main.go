package main

import (
	"fmt"

	"github.com/gopherzz/simpleparserlexer/internal/lexer"
	"github.com/gopherzz/simpleparserlexer/internal/parser"
)

func main() {
	input := "1 + 2 + 3 - 4 + 5"
	lexer := lexer.NewLexer(input)
	tokens := lexer.Tokenize()

	for idx, token := range tokens {
		fmt.Println(idx, token)
	}

	parser := parser.NewParser(tokens)
	expressions := parser.Parse()

	for _, expr := range expressions {
		fmt.Println(expr, "=", expr.Eval())
	}
}
