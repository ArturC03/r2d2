package main

import (
	"fmt"
	"github.com/ArturC03/r2d2/parser"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input := antlr.NewInputStream("import asd from dsa")
	lexer := parser.NewR2DLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewR2DParser(stream)
	p.BuildParseTrees = true
	tree := p.R2D()

	// Correct way to print the parse tree
	fmt.Println(tree.ToStringTree([]string{}, p))
}
