package visitor

import (
	"fmt"
	"r2d2/parser"

	"github.com/antlr4-go/antlr/v4"
)

// R2D2Visitor implementa a interface parser.R2D2Visitor
type R2D2Visitor struct {
	parser.BaseR2D2Visitor
}

// Novo visitor
func NewR2D2Visitor() *R2D2Visitor {
	return &R2D2Visitor{}
}

// VisitChildren implementa a visitação de filhos corretamente
func (v *R2D2Visitor) VisitChildren(node antlr.RuleNode) any {
	var result any

	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		if parseTree, ok := child.(antlr.ParseTree); ok {
			childResult := parseTree.Accept(v)
			// Atualiza o resultado se necessário
			result = childResult
		}
	}

	return result
}

// TODO:
func (v *R2D2Visitor) VisitProgram(ctx *parser.ProgramContext) any {
	fmt.Println("🚀 Visitando Program!")
	return v.VisitChildren(ctx)
}

// TODO:
func (v *R2D2Visitor) VisitDeclaration(ctx *parser.DeclarationContext) any {
	fmt.Println("📄 Visitando Declaration:", ctx.GetText())
	return v.VisitChildren(ctx)
}

// TODO:
func (v *R2D2Visitor) VisitImportDeclaration(ctx *parser.ImportDeclarationContext) any {
	fmt.Println("📥 Import detectado:", ctx.GetText())
	return v.VisitChildren(ctx)
}

// TODO:
func (v *R2D2Visitor) VisitModuleDeclaration(ctx *parser.ModuleDeclarationContext) any {
	fmt.Println("📦 Module detectado:", ctx.GetText())
	return v.VisitChildren(ctx)
}

// TODO:
func (v *R2D2Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) any {
	fmt.Println("📄 Visitando FunctionDeclaration:", ctx.GetText())
	return v.VisitChildren(ctx)
}
