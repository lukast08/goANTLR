package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	parser "github.com/lukast08/ANTLRs/antlr_files"
)

func main() {

	input := antlr.NewInputStream("{'hello': 'world'}")
	lexer := parser.NewJSONLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewJSONParser(tokens)
	fmt.Println(p)
}
