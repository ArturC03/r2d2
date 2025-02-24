package main

import (
	"fmt"

	"github.com/ArturC03/r2d2/parser"
	"github.com/ArturC03/r2d2/visitor"
	r2d2Styles "github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// Criar um stream de entrada
	input := antlr.NewInputStream(`module cookie{
		fn main(){
		loop {
		}
		cookie();
		}
		}`) // Ajuste conforme a gramática

	fmt.Println(r2d2Styles.InfoMessage("Input stream criado."))

	// Criar Lexer
	lexer := parser.NewR2D2Lexer(input)
	if lexer == nil {
		fmt.Println(r2d2Styles.ErrorMessage("Lexer não foi inicializado corretamente!"))
	}
	fmt.Println(r2d2Styles.InfoMessage("Lexer criado."))

	// Criar Token Stream
	stream := antlr.NewCommonTokenStream(lexer, 0)
	fmt.Println(r2d2Styles.InfoMessage("Token stream criado."))

	// Criar Parser
	p := parser.NewR2D2Parser(stream)
	if p == nil {
		fmt.Println(r2d2Styles.ErrorMessage("Parser não foi inicializado corretamente!"))
	}
	fmt.Println(r2d2Styles.InfoMessage("Parser criado."))

	// Construir a árvore de parse
	p.BuildParseTrees = true
	tree := p.Program() // Ajuste conforme sua regra principal
	// Imprimir a árvore sintática
	// fmt.Println(tree.ToStringTree(p.GetRuleNames(), nil))
	// return
	if tree == nil {
		fmt.Println(r2d2Styles.ErrorMessage("Árvore sintática não foi criada corretamente!"))
	}
	fmt.Println(r2d2Styles.InfoMessage("Árvore sintática criada."))

	// Criar e aplicar o Visitor
	v := visitor.NewR2D2Visitor()
	if v == nil {
		fmt.Println(r2d2Styles.ErrorMessage("Visitor não foi inicializado corretamente!"))
	}
	fmt.Println(r2d2Styles.InfoMessage("Visitor criado."))

	// Aplicar o visitor na árvore - esta é a forma correta para ANTLR em Go
	result := tree.Accept(v)

	fmt.Println(r2d2Styles.InfoMessage("Visitor aplicado com sucesso."))

	// Exibir o resultado
	fmt.Println(r2d2Styles.InfoMessage(fmt.Sprintf("Resultado do Visitor: %v", result)))
}
