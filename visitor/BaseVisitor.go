package visitor

import (
	"fmt"
	"os"
	"strings"

	"github.com/ArturC03/r2d2/parser"
	r2d2Styles "github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
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

type Argument struct {
	Name string
	Type string
}

type Function struct {
	Name       string
	Arguments  []Argument // Adicionado para guardar informações dos argumentos
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

	for i := range node.GetChildCount() {
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
		canEscape := findChild(ctx, (*parser.BreakStatementContext)(nil), (*parser.ReturnStatementContext)(nil))

		// No Escape
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

func (v *R2D2Visitor) VisitFunctionCallStatement(ctx *parser.FunctionCallStatementContext) any {
	return v.VisitChildren(ctx)

}

// TODO: Check if the function exists and if it is accessible
func (v *R2D2Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) any {
	loadGlobalFunctions()

	// Nome da função
	funcName := ctx.IDENTIFIER(0).GetText()

	// Verifica se a função existe
	function, exists := v.symbolTable.Modules["global"].Functions[funcName]
	if !exists {
		errorMessage := fmt.Sprintf("/* ERROR: Function '%s' not found */", funcName)
		v.JsCode += errorMessage + "\n"
		fmt.Println(r2d2Styles.ErrorMessage(errorMessage))
		return nil
	}

	// Obtém os argumentos passados na chamada
	var passedArgs []string
	argumentList := ctx.ArgumentList()
	if argumentList != nil {
		for _, arg := range argumentList.AllExpression() {
			passedArgs = append(passedArgs, arg.GetText())
		}
	}

	// Verifica o número de argumentos
	if len(passedArgs) != len(function.Arguments) {
		errorMessage := fmt.Sprintf(
			"/* ERROR: Function '%s' expects %d arguments, but %d were provided */",
			funcName, len(function.Arguments), len(passedArgs),
		)
		v.JsCode += errorMessage + "\n"
		fmt.Println(r2d2Styles.ErrorMessage(errorMessage))
		return nil
	}

	// Verifica os tipos dos argumentos
	for i, passedArg := range passedArgs {
		expectedType := function.Arguments[i].Type

		// Aqui podes implementar lógica para verificar o tipo do argumento
		if !isValidType(passedArg, expectedType) {
			errorMessage := fmt.Sprintf(
				"/* ERROR: Argument %d of function '%s' expects type '%s', but got '%s' */",
				i+1, funcName, expectedType, passedArg,
			)
			v.JsCode += errorMessage + "\n"
			fmt.Println(r2d2Styles.ErrorMessage(errorMessage))
			return nil
		}
	}

	// Monta a chamada correta no código compilado
	v.JsCode += fmt.Sprintf("%s(%s);\n", funcName, strings.Join(passedArgs, ", "))

	return nil
}

func isValidType(value string, expectedType string) bool {
	// Implementa a lógica para verificar tipos (ex.: int, string, etc.)
	// Por exemplo, verificar se "123" é int ou "texto" é string
	return true // Ajusta conforme necessário
}

func (v *R2D2Visitor) VisitVariableDeclarationStatement(ctx *parser.VariableDeclarationContext) any {
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
			v.JsCode += fmt.Sprintf("const %s = %s;", ctx.IDENTIFIER().GetText(), ctx.Expression().GetText())
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

	if !findParent(ctx, (*parser.FunctionDeclarationContext)(nil)) {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Invalid return statement on line %d", ctx.GetStart().GetLine())))
	}

	v.JsCode += "return"
	if ctx.Expression() != nil {
		v.JsCode += " " + ctx.Expression().GetText()
	}
	v.JsCode += ";"
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitCicleControl(ctx *parser.CicleControlContext) any {
	// Verifica se o ciclo de controle (break/continue) está dentro de um loop válido
	if !findParent(ctx,
		(*parser.LoopStatementContext)(nil),
		(*parser.WhileStatementContext)(nil),
		(*parser.ForStatementContext)(nil)) {

		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Invalid %s statement on line %d",
			ctx.GetStart().GetText(), ctx.GetStart().GetLine())))
	}

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitIfStatement(ctx *parser.IfStatementContext) any {
	// Open the initial 'if' block
	v.JsCode += fmt.Sprintf("if (%s) {", ctx.Expression(0).GetText())

	// Visit the main if block (instead of directly visiting, let the block visitor handle it)
	blockCtx := ctx.Block(0)
	blockCtx.Accept(v)

	// Process any 'else if' statements
	for i := 1; i < len(ctx.AllExpression()); i++ {
		v.JsCode += fmt.Sprintf("} else if (%s) {", ctx.Expression(i).GetText())
		ctx.Block(i).Accept(v)
	}

	// Process the 'else' statement if present (checking ELSE token count matches block count)
	elseBlockIndex := len(ctx.AllExpression())
	if len(ctx.AllELSE()) > len(ctx.AllExpression())-1 && elseBlockIndex < len(ctx.AllBlock()) {
		v.JsCode += "} else {"
		ctx.Block(elseBlockIndex).Accept(v)
	}

	// Close the final if-else block
	v.JsCode += "}"

	return nil // Return value is not used for statement visitors
}

func (v *R2D2Visitor) VisitWhileStatement(ctx *parser.WhileStatementContext) any {
	// Obtém a condição do while
	condition := v.Visit(ctx.Expression()).(string)

	// Inicia o bloco while com a condição
	v.JsCode += fmt.Sprintf("while (%s) {", condition)

	// Visita os filhos (o corpo do loop)
	result := v.VisitChildren(ctx)

	// Fecha o bloco while
	v.JsCode += "}"

	return result
}

func (v *R2D2Visitor) VisitForStatement(ctx *parser.ForStatementContext) any {

	v.JsCode += "for ("

	// Debug: imprimir a estrutura do SimpleFor
	simpleFor := ctx.SimpleFor()
	if simpleFor != nil {

		// Inicialização
		if simpleFor.VariableDeclaration() != nil {
			varDecl := simpleFor.VariableDeclaration()

			// Tipo de declaração
			if varDecl.VAR() != nil {
				v.JsCode += "var "
			} else if varDecl.LET() != nil {
				v.JsCode += "let "
			} else if varDecl.CONST() != nil {
				v.JsCode += "const "
			}

			// Identificador
			v.JsCode += varDecl.IDENTIFIER().GetText()

			// Ignorar o tipo em JS (i32)

			// Atribuição
			if varDecl.ASSIGN() != nil && varDecl.Expression() != nil {
				v.JsCode += " = " + varDecl.Expression().GetText()
			}
		} else if len(simpleFor.AllAssignment()) > 0 {
			// Primeira atribuição
			assignment := simpleFor.Assignment(0)
			v.JsCode += assignment.IDENTIFIER().GetText()

			if assignment.AssignmentOperator() != nil {
				v.JsCode += " " + assignment.AssignmentOperator().GetText() + " "
				if assignment.Expression() != nil {
					v.JsCode += assignment.Expression().GetText()
				}
			} else if assignment.INCREMENT() != nil {
				v.JsCode += "++"
			} else if assignment.DECREMENT() != nil {
				v.JsCode += "--"
			}
		}

		// Condição
		v.JsCode += "; "
		if simpleFor.Expression() != nil {
			v.JsCode += simpleFor.Expression().GetText()
		}

		// Atualização
		v.JsCode += "; "
		if len(simpleFor.AllAssignment()) > 1 {
			assignment := simpleFor.Assignment(1)
			identifier := assignment.IDENTIFIER().GetText()

			if assignment.AssignmentOperator() != nil {
				v.JsCode += identifier + " " + assignment.AssignmentOperator().GetText() + " "
				if assignment.Expression() != nil {
					v.JsCode += assignment.Expression().GetText()
				}
			} else if assignment.INCREMENT() != nil {
				v.JsCode += identifier + "++"
			} else if assignment.DECREMENT() != nil {
				v.JsCode += identifier + "--"
			}
		} else if len(simpleFor.AllAssignment()) > 0 {
			// Se não houver uma segunda atribuição explícita, adicionar incremento simples
			v.JsCode += simpleFor.Assignment(0).IDENTIFIER().GetText() + "++"
		}
	}

	v.JsCode += ") {"

	// Processar o bloco
	if ctx.Block() != nil {
		ctx.Block().Accept(v)
	}

	v.JsCode += "}"

	return nil
}

// Função auxiliar para processar atribuições
func (v *R2D2Visitor) VisitAssignment(ctx *parser.AssignmentContext) any {
	// Se tivermos um operador de incremento/decremento pós-fixado
	if ctx.INCREMENT() != nil {
		v.JsCode += ctx.IDENTIFIER().GetText() + "++"
		return nil
	} else if ctx.DECREMENT() != nil {
		v.JsCode += ctx.IDENTIFIER().GetText() + "--"
		return nil
	}

	// Para atribuições regulares
	identifier := ctx.IDENTIFIER().GetText()
	v.JsCode += identifier

	// Mapear os operadores de atribuição
	if ctx.AssignmentOperator() != nil {
		op := ctx.AssignmentOperator().GetText()
		v.JsCode += " " + op + " "
	}

	// Processar a expressão do lado direito
	if ctx.Expression() != nil {
		v.JsCode += ctx.Expression().GetText()
	}

	return nil
}

func (v *R2D2Visitor) VisitArgumentList(ctx *parser.ArgumentListContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitMemberExpression(ctx *parser.MemberExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitExpression(ctx *parser.ExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitComparisonExpression(ctx *parser.ComparisonExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitMemberPart(ctx *parser.MemberPartContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitLiteral(ctx *parser.LiteralContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitSwitchCase(ctx *parser.SwitchCaseContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) any {
	return v.VisitChildren(ctx)
}
