package main

import (
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
	"r2d2/parser"
	"r2d2/visitor"
)

func main() {
	// Criar um stream de entrada
	input := antlr.NewInputStream(`module cookie{fn main(){}}`) // Ajuste conforme a gramática
	log.Println("✅ Input stream criado.")

	// Criar Lexer
	lexer := parser.NewR2D2Lexer(input)
	if lexer == nil {
		log.Fatal("❌ Erro: Lexer não foi inicializado corretamente!")
	}
	log.Println("✅ Lexer criado.")

	// Criar Token Stream
	stream := antlr.NewCommonTokenStream(lexer, 0)
	log.Println("✅ Token stream criado.")

	// Criar Parser
	p := parser.NewR2D2Parser(stream)
	if p == nil {
		log.Fatal("❌ Erro: Parser não foi inicializado corretamente!")
	}
	log.Println("✅ Parser criado.")

	// Construir a árvore de parse
	p.BuildParseTrees = true
	tree := p.Program() // Ajuste conforme sua regra principal
	if tree == nil {
		log.Fatal("❌ Erro: Árvore sintática não foi criada corretamente!")
	}
	log.Println("✅ Árvore sintática criada.")

	// Criar e aplicar o Visitor
	v := visitor.NewR2D2Visitor()
	if v == nil {
		log.Fatal("❌ Erro: Visitor não foi inicializado corretamente!")
	}
	log.Println("✅ Visitor criado.")

	// Aplicar o visitor na árvore - esta é a forma correta para ANTLR em Go
	result := tree.Accept(v)

	log.Println("✅ Visitor aplicado com sucesso.")

	// Exibir o resultado
	fmt.Println("🚀 Resultado do Visitor:", result)
}
