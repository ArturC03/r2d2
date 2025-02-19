package visitor

import (
	"fmt"
	"github.com/ArturC03/r2d2/parser"
)

// Criar o struct do Visitor
type MyVisitor struct {
	*parser.BaseR2D2Visitor
}

// Inicializador do visitor
func NewMyVisitor() *MyVisitor {
	return &MyVisitor{}
}

// Exemplo: Visitar o nó principal da gramática
func (v *MyVisitor) VisitR2D2(ctx *parser.ProgramContext) interface{} {
	fmt.Println("🚀 Visitando nó principal: R2D2")
	return v.VisitChildren(ctx) // Continua visitando os filhos
}
