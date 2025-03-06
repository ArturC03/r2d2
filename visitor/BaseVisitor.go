package visitor

import (
	"fmt"
	"github.com/ArturC03/r2d2/parser"
	r2d2Styles "github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"strings"
)

type Variable struct {
	Name       string
	Value      any
	Type       string
	isExported bool
}

type Global struct {
	Name  string
	Value any
	Type  string
}

type Function struct {
	Name       string
	Variables  map[string]Variable
	Functions  map[string]Function
	isExported bool
	isPseudo   bool
}

type Module struct {
	Name      string
	Functions map[string]Function
	Variables map[string]Variable
	Types     map[string]any
}

type Interface struct {
	Name      string
	Functions map[string]Function
}

type SymbolTable struct {
	Modules    map[string]Module
	Interfaces map[string]Interface
	Globals    map[string]Global
}

type R2D2Visitor struct {
	parser.BaseR2D2Visitor
	symbolTable SymbolTable
	JsCode      string
}

func NewR2D2Visitor() *R2D2Visitor {
	return &R2D2Visitor{
		symbolTable: SymbolTable{
			Modules:    make(map[string]Module),
			Interfaces: make(map[string]Interface),
			Globals:    make(map[string]Global),
		},
		JsCode: "",
	}
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
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitDeclaration(ctx *parser.DeclarationContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitImportDeclaration(ctx *parser.ImportDeclarationContext) any {

	// Has String Literal
	if ctx.STRING_LITERAL() == nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("File path not found on line %d", ctx.GetStart().GetLine())))
		return v.VisitChildren(ctx)
	}

	path := ctx.STRING_LITERAL().GetText() // path of file with no quotes

	// Empty String
	if path == "\"\"" {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Empty file path on line %d", ctx.GetStart().GetLine())))

	} else {
		justPath := strings.Trim(path, "\"")
		_, err := os.Stat(justPath)
		if err != nil {
			fmt.Println(r2d2Styles.ErrorMessage("File not found on path " + justPath))
		}
	}

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitGlobalDeclaration(ctx *parser.GlobalDeclarationContext) any {

	v.symbolTable.Globals[ctx.IDENTIFIER().GetText()] = Global{
		Value: ctx.Expression().GetText(),
		Type:  ctx.TypeExpression().GetText(),
	}

	jsCode := fmt.Sprintf("const %s = %s;", ctx.IDENTIFIER().GetText(), ctx.Expression().GetText())

	v.JsCode += jsCode

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitModuleDeclaration(ctx *parser.ModuleDeclarationContext) any {
	moduleName := ctx.IDENTIFIER(0).GetText()

	// Create Module
	if _, exists := v.symbolTable.Modules[moduleName]; !exists {
		v.symbolTable.Modules[moduleName] = Module{
			Functions: make(map[string]Function),
			Variables: make(map[string]Variable),
			Types:     make(map[string]any),
		}
		// fmt.Println(r2d2Styles.InfoMessage("Módulo " + moduleName + " criado"))
	}

	// Function Declaration
	for _, child := range ctx.GetChildren() {

		// Function Declaration
		if funcDecl, ok := child.(*parser.FunctionDeclarationContext); ok {
			funcName := funcDecl.IDENTIFIER().GetText()

			function := Function{
				Name:       funcName,
				Variables:  make(map[string]Variable),
				Functions:  make(map[string]Function),
				isExported: isExported(funcDecl),
				isPseudo:   isPseudo(funcDecl),
			}

			// Add Function to Module
			v.symbolTable.Modules[moduleName].Functions[funcName] = function
		}

		// Variable Declaration
		if varDecl, ok := child.(*parser.VariableDeclarationContext); ok {
			varName := varDecl.IDENTIFIER().GetText()

			variable := Variable{
				Name:       varName,
				Value:      nil,
				Type:       varDecl.TypeExpression().GetText(),
				isExported: isExported(varDecl),
			}

			// Variable Declaration with Assignment
			if varDecl.ASSIGN() != nil {
				variable.Value = varDecl.Expression().GetText()
			} else {
			}
			v.symbolTable.Modules[moduleName].Variables[varName] = variable
		}

		// Type Declaration
		if typeDecl, ok := child.(*parser.TypeDeclarationContext); ok {
			typeName := typeDecl.IDENTIFIER().GetText()
			v.symbolTable.Modules[moduleName].Types[typeName] = typeDecl
			// fmt.Println(r2d2Styles.InfoMessage("Tipo " + typeName + " criado no módulo " + moduleName))
		}
	}

	// Start Module
	v.JsCode += fmt.Sprintf("const %s = (function () {\n", moduleName)

	result := v.VisitChildren(ctx)

	moduleExports := v.symbolTable.Modules[moduleName].Exports() // Exported Assets

	// End Module
	v.JsCode += fmt.Sprintf("return {%s}; })();\n", strings.Join(moduleExports, ", "))

	return result
}

func (v *R2D2Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) any {

	// Pseudo
	if ctx.PSEUDO() != nil {
	} else {
		// fmt.Println(r2d2Styles.InfoMessage("Function declaration: " + ctx.GetText()))
	}

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitBlock(ctx *parser.BlockContext) any {
	// fmt.Println(r2d2Styles.InfoMessage("Visiting block: " + ctx.GetText()))

	// Function Declaration
	if parentFuncDecl, ok := ctx.GetParent().(*parser.FunctionDeclarationContext); ok {

		// Pseudo
		if parentFuncDecl.PSEUDO() != nil {
			fmt.Println(r2d2Styles.InfoMessage("Found block inside a pseudo function: " + ctx.GetText()))

			for _, child := range ctx.GetChildren() {

				// Statement
				if stmtCtx, ok := child.(*parser.StatementContext); ok {

					// Not FunctionCall
					if _, ok := stmtCtx.GetChild(0).(*parser.FunctionCallStatementContext); !ok {
						line := stmtCtx.GetStart().GetLine()
						fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Line %d: statement %s not allowed in a pseudo function", line, stmtCtx.GetStart().GetText())))
					} else {
						// FunctionCall
						// v.VisitChildren(stmtCtx)
					}
				}
			}
		} else {
			for _, child := range ctx.GetChildren() {

				// Statement
				if _, ok := child.(*parser.StatementContext); ok {
					// v.VisitChildren(stmtCtx)
				}
			}
		}
	}

	// Loop
	if parentLoop, ok := ctx.GetParent().(*parser.LoopStatementContext); ok {

		canEscape := false
		for _, child := range ctx.GetChildren() {

			// Statement
			if stmtCtx, ok := child.(*parser.StatementContext); ok {

				// Break
				if loopCtrl, ok := stmtCtx.GetChild(0).(*parser.CicleControlContext); ok && loopCtrl.BREAK() != nil {
					canEscape = true
					break
				}

				// Return
				if _, ok := stmtCtx.GetChild(0).(*parser.ReturnStatementContext); ok {
					canEscape = true
					break
				}
			}
		}
		// No Excape
		if !canEscape {
			line := parentLoop.GetStart().GetLine()
			fmt.Println(r2d2Styles.WarningMessage(fmt.Sprintf("Loop on line %d has no escape!", line)))
		}
	}

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitLoopStatement(ctx *parser.LoopStatementContext) any {
	v.JsCode += fmt.Sprintf("while (true){")

	result := v.VisitChildren(ctx)

	v.JsCode += "}"

	return result

	// v.JsCode += fmt.Sprintf("const %s = (function () {\n", moduleName)
	//
	// result := v.VisitChildren(ctx)
	//
	// moduleExports := v.symbolTable.Modules[moduleName].Exports() // Exported Assets
	//
	// // End Module
	// v.JsCode += fmt.Sprintf("return {%s}; })();\n", strings.Join(moduleExports, ", "))
	//
	// return result
}

// TODO: Check if the function exists and if it is accessible
func (v *R2D2Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) any {

	// if ctx.DOT() != nil {
	// 	if v.symbolTable.Modules[ctx.IDENTIFIER(0).GetText()].Functions[ctx.IDENTIFIER(1).GetText()] != nil {
	//
	// 	}

	// }

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) any {
	// Parent is Module
	if _, ok := ctx.GetParent().(*parser.ModuleDeclarationContext); !ok {

		// Export
		if ctx.EXPORT() != nil {
			fmt.Println(r2d2Styles.ErrorMessage("Cannot export non-global variables"))
		}
	}

	// Temporary Variable
	if ctx.LET() != nil {
		v.JsCode += fmt.Sprintf("let %s", ctx.IDENTIFIER().GetText())

		// Constant
	} else if ctx.CONST() != nil {
		if ctx.ASSIGN() == nil {
			fmt.Println(r2d2Styles.ErrorMessage("Const variable must be assigned a value"))
		} else {
			v.JsCode += fmt.Sprintf("const %s", ctx.IDENTIFIER().GetText(), ctx.Expression().GetText())
		}

		// Variable
	} else if ctx.VAR() != nil {
		v.JsCode += fmt.Sprintf("var %s", ctx.IDENTIFIER().GetText())
	}

	if ctx.ASSIGN() != nil {
		v.JsCode += fmt.Sprintf(" = %s;", ctx.Expression().GetText())
	} else {
		v.JsCode += ";"
	}
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitStatement(ctx *parser.StatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) any {
	v.JsCode += "return"
	if ctx.Expression() != nil {
		v.JsCode += " " + ctx.Expression().GetText()
	}
	v.JsCode += ";"
	return v.VisitChildren(ctx)
}

// TODO: Apenas se o pai for um loop o cicle Control é válido
// func (v *R2D2Visitor) VisitCicleControl(ctx *parser.CicleControlContext) any {
//
// 	// if parentCicle, ok := ctx.GetParent().(*parser.LoopStatementContext); ok {
// 	// if parentCicle, ok := ctx.GetParent().(*parser.WhileStatementContext); ok {
// 	// if parentCicle, ok := ctx.GetParent().(*parser.ForStatementContext); ok {
// 	//
//
// 	return v.VisitChildren(ctx)
// }
