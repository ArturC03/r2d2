package visitor

import (
	"fmt"
	"github.com/ArturC03/r2d2/parser"
	r2d2Styles "github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
)

type R2D2Visitor struct {
	parser.BaseR2D2Visitor
}

func NewR2D2Visitor() *R2D2Visitor {
	return &R2D2Visitor{}
}

func (v *R2D2Visitor) VisitChildren(node antlr.RuleNode) any {
	var result any

	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		if parseTree, ok := child.(antlr.ParseTree); ok {
			childResult := parseTree.Accept(v)
			// Update result if necessary
			result = childResult
		}
	}

	return result
}

func (v *R2D2Visitor) VisitProgram(ctx *parser.ProgramContext) any {
	fmt.Println(r2d2Styles.InfoMessage("Visitando Program!"))
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitDeclaration(ctx *parser.DeclarationContext) any {
	fmt.Println(r2d2Styles.InfoMessage("Visitando Declaration: " + ctx.GetText()))
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitImportDeclaration(ctx *parser.ImportDeclarationContext) any {
	fmt.Println(r2d2Styles.InfoMessage("Import detectado: " + ctx.GetText()))
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitModuleDeclaration(ctx *parser.ModuleDeclarationContext) any {
	fmt.Println(r2d2Styles.InfoMessage("Module detectado: " + ctx.GetText()))
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) any {
	// Pseudo function
	if ctx.PSEUDO() != nil {
		// fmt.Println(r2d2Styles.WarningMessage("Pseudo function declaration detected: " + ctx.GetText()))

	} else {
		// Directly call InfoMessage without .Render
		// fmt.Println(r2d2Styles.InfoMessage("Function declaration: " + ctx.GetText()))
	}
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitBlock(ctx *parser.BlockContext) any {

	// Function
	if parentFuncDecl, ok := ctx.GetParent().(*parser.FunctionDeclarationContext); ok {

		// Pseudo function
		if parentFuncDecl.PSEUDO() != nil {
			// fmt.Println(r2d2Styles.InfoMessage("Encontrado bloco dentro de uma função pseudo: " + ctx.GetText()))

			// For each statement inside the block
			for _, child := range ctx.GetChildren() {
				if stmtCtx, ok := child.(*parser.StatementContext); ok {

					// If it is a function call
					if functionCallStmt, ok := stmtCtx.GetChild(0).(*parser.FunctionCallStatementContext); ok {
						// fmt.Println(r2d2Styles.InfoMessage("Encontrado FunctionCallStatement: " + functionCallStmt.GetText()))
					} else {
						// Invalid statement inside pseudo function
						fmt.Println(r2d2Styles.ErrorMessage("Encontrado statement não permitido: " + stmtCtx.GetText()))
						// return fmt.Errorf("função pseudo só pode conter chamadas de função")
					}
				}
			}
		}
	}

	return v.VisitChildren(ctx)
}
