package main

import (
	"fmt"
	"github.com/ArturC03/r2d2/parser"
	"github.com/ArturC03/r2d2/visitor"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// Código de entrada para ser analisado
	input := antlr.NewInputStream("import asd from dsa")

	// Criar Lexer, Token Stream e Parser
	lexer := parser.NewR2D2Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewR2D2Parser(stream)
	p.BuildParseTrees = true

	// Criar árvore sintática
	tree := p.R2D2() // 🛠️ Certifica-te de usar a regra correta!

	// Criar e aplicar o visitor
	v := visitor.NewMyVisitor()
	result := v.Visit(tree)

	fmt.Println("🚀 Resultado do Visitor:", result)
}
