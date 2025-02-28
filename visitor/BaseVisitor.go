package visitor

import (
	"fmt"
	"github.com/ArturC03/r2d2/parser"
	r2d2Styles "github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"strings"
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
	// filePath := ""
	path := ctx.STRING_LITERAL().GetText()

	if ctx.STRING_LITERAL() == nil || path == "\"\"" {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("File path not found on line %d", ctx.GetStart().GetLine())))
	} else if _, err := os.Stat(ctx.STRING_LITERAL().GetText()); err != nil {
		justPath := strings.Trim(path, "\"")
		fmt.Println(r2d2Styles.ErrorMessage("File not found on path " + justPath))
	}

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitModuleDeclaration(ctx *parser.ModuleDeclarationContext) any {
	fmt.Println(r2d2Styles.InfoMessage("Module detectado: " + ctx.GetText()))
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) any {
	if ctx.PSEUDO() != nil {
		fmt.Println(r2d2Styles.WarningMessage("Pseudo function declaration detected: " + ctx.GetText()))
	} else {
		fmt.Println(r2d2Styles.InfoMessage("Function declaration: " + ctx.GetText()))
	}
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitBlock(ctx *parser.BlockContext) any {
	fmt.Println(r2d2Styles.InfoMessage("Visiting block: " + ctx.GetText()))

	// Verifica se o bloco está dentro de uma função pseudo
	if parentFuncDecl, ok := ctx.GetParent().(*parser.FunctionDeclarationContext); ok {

		if parentFuncDecl.PSEUDO() != nil {
			fmt.Println(r2d2Styles.InfoMessage("Found block inside a pseudo function: " + ctx.GetText()))

			// Iterar sobre os filhos do bloco e verificar declarações inválidas
			for _, child := range ctx.GetChildren() {
				if stmtCtx, ok := child.(*parser.StatementContext); ok {
					// Verifica se o primeiro filho não é uma chamada de função
					if _, ok := stmtCtx.GetChild(0).(*parser.FunctionCallStatementContext); !ok {
						line := stmtCtx.GetStart().GetLine()
						fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Line %d: statement %s not allowed in a pseudo function", line, stmtCtx.GetStart().GetText())))
					} else {
						// Se for uma chamada de função válida, continue visitando
						v.VisitChildren(stmtCtx)
					}
				}
			}
		} else {
			for _, child := range ctx.GetChildren() {
				if stmtCtx, ok := child.(*parser.StatementContext); ok {
					v.VisitChildren(stmtCtx)
				}
			}
		}
	}

	// Verifica se o bloco está dentro de um loop
	if parentLoop, ok := ctx.GetParent().(*parser.LoopStatementContext); ok {
		fmt.Println(r2d2Styles.InfoMessage("Loop block detected: " + ctx.GetText()))

		canEscape := false
		for _, child := range ctx.GetChildren() {
			if stmtCtx, ok := child.(*parser.StatementContext); ok {
				// Verifica controle de loop (BREAK)
				if loopCtrl, ok := stmtCtx.GetChild(0).(*parser.LoopControlContext); ok && loopCtrl.BREAK() != nil {
					canEscape = true
					break
				}
				// Verifica se há uma declaração de retorno
				if _, ok := stmtCtx.GetChild(0).(*parser.ReturnStatementContext); ok {
					canEscape = true
					break
				}
			}
		}
		if !canEscape {
			line := parentLoop.GetStart().GetLine()
			fmt.Println(r2d2Styles.WarningMessage(fmt.Sprintf("Loop on line %d has no escape!", line)))
		}
	}

	// **Aqui chamamos VisitChildren após todas as verificações, sem mudar sua estrutura**
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitLoopStatement(ctx *parser.LoopStatementContext) any {
	fmt.Println(r2d2Styles.InfoMessage("Loop detectado: " + ctx.GetText()))
	return v.VisitChildren(ctx)
}
func (v *R2D2Visitor) VisitFunctionCallStatement(ctx *parser.FunctionCallStatementContext) any {
	fmt.Println(r2d2Styles.InfoMessage("FunctionCall detectado: " + ctx.GetText()))
	return v.VisitChildren(ctx)
}
