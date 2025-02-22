package visitor

import (
	"fmt"
	"github.com/ArturC03/r2d2/parser"
	"github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
)

// R2D2Visitor implements the parser.R2D2Visitor interface
type R2D2Visitor struct {
	parser.BaseR2D2Visitor
}

// NewR2D2Visitor creates a new R2D2Visitor
func NewR2D2Visitor() *R2D2Visitor {
	return &R2D2Visitor{}
}

// VisitChildren correctly implements the visitation of child nodes
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

// VisitProgram handles the Program node visitation
func (v *R2D2Visitor) VisitProgram(ctx *parser.ProgramContext) any {
	fmt.Println(r2d2Styles.InfoMessage.Render("🚀 Visitando Program!"))
	return v.VisitChildren(ctx)
}

// VisitDeclaration handles the Declaration node visitation
func (v *R2D2Visitor) VisitDeclaration(ctx *parser.DeclarationContext) any {
	fmt.Println(r2d2Styles.InfoMessage.Render("📄 Visitando Declaration: " + ctx.GetText()))
	return v.VisitChildren(ctx)
}

// VisitImportDeclaration handles the ImportDeclaration node visitation
func (v *R2D2Visitor) VisitImportDeclaration(ctx *parser.ImportDeclarationContext) any {
	fmt.Println(r2d2Styles.InfoMessage.Render("📥 Import detectado: " + ctx.GetText()))
	return v.VisitChildren(ctx)
}

// VisitModuleDeclaration handles the ModuleDeclaration node visitation
func (v *R2D2Visitor) VisitModuleDeclaration(ctx *parser.ModuleDeclarationContext) any {
	fmt.Println(r2d2Styles.InfoMessage.Render("📦 Module detectado: " + ctx.GetText()))
	return v.VisitChildren(ctx)
}

// VisitFunctionDeclaration handles the FunctionDeclaration node visitation
func (v *R2D2Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) any {
	// Example of adding validation, you can expand on this
	if ctx.PSEUDO() != nil {
		fmt.Println(r2d2Styles.WarningMessage.Render("⚠️ Pseudo function declaration detected: " + ctx.GetText()))
	} else {
		fmt.Println(r2d2Styles.InfoMessage.Render("🔧 Function declaration: " + ctx.GetText()))
	}
	return v.VisitChildren(ctx)
}

// VisitBlock handles the Block node visitation, adding pseudo function validation
func (v *R2D2Visitor) VisitBlock(ctx *parser.BlockContext) any {
	// Verifies if the block belongs to a pseudo function
	if parentFuncDecl, ok := ctx.GetParent().(*parser.FunctionDeclarationContext); ok {
		if parentFuncDecl.PSEUDO() != nil {
			// If the parent function is a pseudo function, we check its contents
			fmt.Println(r2d2Styles.InfoMessage.Render("📄 Encontrado bloco dentro de uma função pseudo: " + ctx.GetText()))

			// Iterate over children to check statements
			for _, child := range ctx.GetChildren() {
				if stmtCtx, ok := child.(*parser.StatementContext); ok {
					if functionCallStmt, ok := stmtCtx.GetChild(0).(*parser.FunctionCallStatementContext); ok {
						// Function calls are valid inside pseudo functions
						fmt.Println(r2d2Styles.InfoMessage.Render("✅ Encontrado FunctionCallStatement: " + functionCallStmt.GetText()))
					} else {
						// Invalid statement inside pseudo function
						fmt.Println(r2d2Styles.ErrorMessage.Render("❌ Encontrado statement não permitido: " + stmtCtx.GetText()))
						return fmt.Errorf("função pseudo só pode conter chamadas de função")
					}
				}
			}
		}
	}

	return v.VisitChildren(ctx)
}
