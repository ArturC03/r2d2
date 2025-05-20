package visitor

import (
	"fmt"
	"os"
	"strings"

	"github.com/ArturC03/r2d2/parser"
	"github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
)

// formatErrorMessage creates a standardized error message with line number
func formatErrorMessage(message string, line int) string {
	return fmt.Sprintf("%s on line %s", message, r2d2Styles.Bold(fmt.Sprintf("%d", line)))
}

// formatWarningMessage creates a standardized warning message with line number
func formatWarningMessage(message string, line int) string {
	return fmt.Sprintf("%s on line %s", message, r2d2Styles.Bold(fmt.Sprintf("%d", line)))
}

// formatErrorMessageNoLine creates a standardized error message without line number
func formatErrorMessageNoLine(message string) string {
	return message
}

// formatWarningMessageNoLine creates a standardized warning message without line number
func formatWarningMessageNoLine(message string) string {
	return message
}

// formatFileErrorMessage creates a standardized error message for file operations
func formatFileErrorMessage(message string, filePath string) string {
	return fmt.Sprintf("%s: %s", message, r2d2Styles.Bold(filePath))
}

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
	Arguments  map[string]Argument
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
	symbolTable      SymbolTable
	JsCode           string
	currentModule    Module
	currentFunction  Function
	currentInterface Interface
}

// Exports returns a list of exported function and variable names
func (m Module) Exports() []string {
	var exports []string

	// Add exported functions
	for name, function := range m.Functions {
		if function.isExported {
			exports = append(exports, name)
		}
	}

	// Add exported variables
	for name, variable := range m.Variables {
		if variable.isExported {
			exports = append(exports, name)
		}
	}

	return exports
}

func NewR2D2Visitor() *R2D2Visitor {
	return &R2D2Visitor{
		symbolTable: SymbolTable{
			Modules:    make(map[string]Module),
			Interfaces: make(map[string]Interface),
			Globals:    make(map[string]Global),
		},
		currentModule: Module{
			Name:      "",
			Functions: make(map[string]Function),
			Variables: make(map[string]Variable),
			Types:     make(map[string]any),
		},
		currentFunction:  Function{},
		currentInterface: Interface{},
		JsCode:           "",
	}
}

// Helper functions
func isExported(node any) bool {
	// Check if the node has an EXPORT token
	switch n := node.(type) {
	case *parser.FunctionDeclarationContext:
		return n.EXPORT() != nil
	case *parser.VariableDeclarationContext:
		return n.EXPORT() != nil
	case *parser.TypeDeclarationContext:
		return n.EXPORT() != nil
	}
	return false
}

func isPseudo(node *parser.FunctionDeclarationContext) bool {
	return node.PSEUDO() != nil
}

func findChild(parent antlr.RuleContext, types ...any) bool {
	for i := range parent.GetChildCount() {
		child := parent.GetChild(i)

		// Check if child matches any of the target types
		for _, t := range types {
			if ctx, ok := child.(antlr.RuleContext); ok {
				switch t.(type) {
				case *parser.BreakStatementContext:
					if _, ok := ctx.(*parser.BreakStatementContext); ok {
						return true
					}
				case *parser.ReturnStatementContext:
					if _, ok := ctx.(*parser.ReturnStatementContext); ok {
						return true
					}
					// Add other cases as needed
				}
			}
		}

		// Recursively check children
		if ctx, ok := child.(antlr.RuleContext); ok {
			if findChild(ctx, types...) {
				return true
			}
		}
	}
	return false
}

func findParent(node antlr.RuleContext, types ...any) bool {
	parent := node.GetParent()
	if parent == nil {
		return false
	}

	// Check if parent matches any of the target types
	for _, t := range types {
		switch t.(type) {
		case *parser.LoopStatementContext:
			if _, ok := parent.(*parser.LoopStatementContext); ok {
				return true
			}
		case *parser.ForStatementContext:
			if _, ok := parent.(*parser.ForStatementContext); ok {
				return true
			}
		case *parser.WhileStatementContext:
			if _, ok := parent.(*parser.WhileStatementContext); ok {
				return true
			}
			// Add other cases as needed
		}
	}

	// Recursively check parent's parent
	return findParent(parent.(antlr.RuleContext), types...)
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

// TODO: Add support for importing files
func (v *R2D2Visitor) VisitImportDeclaration(ctx *parser.ImportDeclarationContext) any {
	// Check if the file path is present
	if ctx.STRING_LITERAL() == nil {
		fmt.Println(r2d2Styles.ErrorMessage(formatFileErrorMessage("File path not found", fmt.Sprintf("%d", ctx.GetStart().GetLine()))))
		return nil
	}

	// Extract and clean the file path
	rawPath := ctx.STRING_LITERAL().GetText()
	if rawPath == "\"\"" {
		fmt.Println(r2d2Styles.ErrorMessage(formatFileErrorMessage("Empty file path", fmt.Sprintf("%d", ctx.GetStart().GetLine()))))
		return nil
	}

	justPath := strings.Trim(rawPath, "\"")

	// Check if the file exists
	if _, err := os.Stat(justPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println(r2d2Styles.ErrorMessage(formatFileErrorMessage("File not found on path", justPath)))
		} else {
			fmt.Println(r2d2Styles.ErrorMessage(formatFileErrorMessage("Error checking file", err.Error())))
		}
		return nil
	}

	// Read the content of the file
	content, err := os.ReadFile(justPath)
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(formatFileErrorMessage("Failed to read file", justPath)))
		return nil
	}

	// Parse the content as if it were part of the same code
	input := antlr.NewInputStream(string(content))
	lexer := parser.NewR2D2Lexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewR2D2Parser(tokens)
	el := parser.NewR2D2ErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	tree := p.Program()

	// Create a sub-visitor to process the imported content
	sub := NewR2D2Visitor()
	tree.Accept(sub)

	// Merge the symbol tables
	for name, module := range sub.symbolTable.Modules {
		v.symbolTable.Modules[name] = module
	}
	for name, iface := range sub.symbolTable.Interfaces {
		v.symbolTable.Interfaces[name] = iface
	}
	for name, global := range sub.symbolTable.Globals {
		v.symbolTable.Globals[name] = global
	}

	// Prepend the imported code
	v.JsCode = sub.JsCode + v.JsCode

	return nil
}

func (v *R2D2Visitor) VisitGlobalDeclaration(ctx *parser.GlobalDeclarationContext) any {
	if ctx.IDENTIFIER() == nil {
		fmt.Println("Error: IDENTIFIER is nil in GlobalDeclaration")
		return nil
	}

	if ctx.Expression() == nil {
		fmt.Println("Error: Expression is nil in GlobalDeclaration")
		return nil
	}

	if ctx.TypeExpression() == nil {
		fmt.Println("Error: TypeExpression is nil in GlobalDeclaration")
		return nil
	}

	globalName := ctx.IDENTIFIER().GetText()

	// Guarda no symbolTable
	v.symbolTable.Globals[globalName] = Global{
		Name:  globalName,
		Value: "",
		Type:  ctx.TypeExpression().GetText(),
	}

	// Gera o JS
	v.JsCode += fmt.Sprintf("const %s = ", globalName)
	ctx.Expression().Accept(v)
	v.JsCode += ";"

	return nil
}

func (v *R2D2Visitor) VisitModuleDeclaration(ctx *parser.ModuleDeclarationContext) any {
	// Garantir que a tabela de símbolos global seja inicializada
	if v.symbolTable.Globals == nil {
		v.symbolTable.Globals = make(map[string]Global)
	}

	if v.symbolTable.Modules == nil {
		v.symbolTable.Modules = make(map[string]Module)
	}

	// Verificação para garantir que o módulo tenha um identificador
	if ctx.IDENTIFIER(0) == nil || len(ctx.AllIDENTIFIER()) == 0 {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Module declaration without identifier")))
		return nil
	}

	moduleName := ctx.IDENTIFIER(0).GetText()

	// Criar o módulo se não existir
	if _, exists := v.symbolTable.Modules[moduleName]; !exists {
		newModule := Module{
			Name:      moduleName,
			Functions: make(map[string]Function),
			Variables: make(map[string]Variable),
			Types:     make(map[string]any),
		}
		v.symbolTable.Modules[moduleName] = newModule
		v.currentModule = newModule
		defer func() {
			v.symbolTable.Modules[v.currentModule.Name] = v.currentModule
		}()
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine(fmt.Sprintf("Module '%s' already exists", moduleName))))
		return nil
	}

	// Processar o conteúdo do módulo para popular a tabela de símbolos
	for _, child := range ctx.GetChildren() {
		// Declaração de função
		if funcDecl, ok := child.(*parser.FunctionDeclarationContext); ok {
			if funcDecl.IDENTIFIER() == nil {
				continue
			}

			funcName := funcDecl.IDENTIFIER().GetText()

			// Parse dos argumentos da função
			arguments := make(map[string]Argument)
			if funcDecl.ParameterList() != nil {
				for _, param := range funcDecl.ParameterList().AllParameter() {
					if param.IDENTIFIER() != nil && param.TypeExpression() != nil {
						paramName := param.IDENTIFIER().GetText()
						arguments[paramName] = Argument{
							Name: paramName,
							Type: param.TypeExpression().GetText(),
						}
					}
				}
			}

			function := Function{
				Name:       funcName,
				Arguments:  arguments,
				Variables:  make(map[string]Variable),
				Functions:  make(map[string]Function),
				isExported: isExported(funcDecl),
				isPseudo:   isPseudo(funcDecl),
			}

			// Armazenar a função no módulo
			module := v.symbolTable.Modules[moduleName]
			module.Functions[funcName] = function
			v.symbolTable.Modules[moduleName] = module
		}

		// Declaração de variável
		if varDecl, ok := child.(*parser.VariableDeclarationContext); ok {
			if varDecl.IDENTIFIER() == nil || varDecl.TypeExpression() == nil {
				continue
			}

			varName := varDecl.IDENTIFIER().GetText()

			variable := Variable{
				Name:       varName,
				Type:       varDecl.TypeExpression().GetText(),
				isExported: isExported(varDecl),
			}

			// Declaração de variável sem atribuição de expressão
			if varDecl.ASSIGN() != nil && varDecl.Expression() != nil {
				// O valor da variável não será extraído, mantendo o comportamento que você deseja
				// Continue sem pegar o texto da expressão
				// variável.Value = varDecl.Expression().GetText()
				expr := varDecl.Expression()
				variable.Value = expr.GetText()
			} else {
				variable.Value = ""
			}

			// Armazenar a variável no módulo
			module := v.symbolTable.Modules[moduleName]
			module.Variables[varName] = variable
			v.symbolTable.Modules[moduleName] = module
		}

		// Declaração de tipo
		if typeDecl, ok := child.(*parser.TypeDeclarationContext); ok {
			if typeDecl.IDENTIFIER() == nil {
				continue
			}

			typeName := typeDecl.IDENTIFIER().GetText()

			// Armazenar o tipo no módulo
			module := v.symbolTable.Modules[moduleName]
			module.Types[typeName] = typeDecl
			v.symbolTable.Modules[moduleName] = module
		}
	}

	// Iniciar o módulo no código JS
	v.JsCode += fmt.Sprintf("const %s = (function () {", moduleName)

	// Visitar os filhos para processar o conteúdo do módulo
	result := v.VisitChildren(ctx)

	// Obter os itens exportados para a instrução de retorno do módulo
	moduleExports := v.symbolTable.Modules[moduleName].Exports()

	// Finalizar o módulo com a instrução de retorno
	v.JsCode += fmt.Sprintf("return {%s}; })();", strings.Join(moduleExports, ", "))

	// Verificar se o módulo exporta uma função main e executá-la
	module := v.symbolTable.Modules[moduleName]
	if mainFunc, exists := module.Functions["main"]; exists && mainFunc.isExported {
		v.JsCode += fmt.Sprintf("%s.main();", moduleName)
	}

	return result
}

func (v *R2D2Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) any {

	// Skip if no identifier
	if ctx.IDENTIFIER() == nil {
		return v.VisitChildren(ctx)
	}
	funcName := ctx.IDENTIFIER().GetText()
	v.currentFunction = Function{
		Name:       funcName,
		Arguments:  make(map[string]Argument),
		Variables:  make(map[string]Variable),
		Functions:  make(map[string]Function),
		isExported: isExported(ctx),
		isPseudo:   isPseudo(ctx),
	}
	defer func() {
		v.symbolTable.Modules[v.currentModule.Name].Functions[v.currentFunction.Name] = v.currentFunction
	}()
	// moduleName := v.currentModule.Name

	// Generate function signature for JS
	v.JsCode += fmt.Sprintf("function %s(", funcName)

	// Add parameters
	if ctx.ParameterList() != nil {
		paramNames := []string{}

		for _, param := range ctx.ParameterList().AllParameter() {
			if param.IDENTIFIER() != nil {
				paramName := param.IDENTIFIER().GetText()
				paramNames = append(paramNames, paramName)

				// Create an Argument object
				argument := Argument{
					Name: paramName,
					Type: param.TypeExpression().GetText(),
				}

				// Add as both an argument and a variable
				v.currentFunction.Arguments[paramName] = argument

				// Also add as a variable in the function's scope
				v.currentFunction.Variables[paramName] = Variable{
					Name:       paramName,
					Type:       param.TypeExpression().GetText(),
					isExported: false,
				}
			}
		}

		v.JsCode += strings.Join(paramNames, ", ")
	}

	v.JsCode += ") {"

	// Store any local variables defined in the function's block
	if ctx.Block() != nil {
		// Visit the function body
		ctx.Block().Accept(v)
	}

	v.JsCode += "}"

	v.symbolTable.Modules[v.currentModule.Name].Functions[v.currentFunction.Name] = v.currentFunction

	return nil
}

func (v *R2D2Visitor) VisitBlock(ctx *parser.BlockContext) any {
	// Function Declaration
	if parentFuncDecl, ok := ctx.GetParent().(*parser.FunctionDeclarationContext); ok {
		// Pseudo
		if parentFuncDecl.PSEUDO() != nil {
			for _, child := range ctx.GetChildren() {
				// Statement
				if stmtCtx, ok := child.(*parser.StatementContext); ok {
					// Not FunctionCall
					if _, ok := stmtCtx.GetChild(0).(*parser.FunctionCallStatementContext); !ok {
						line := stmtCtx.GetStart().GetLine()
						fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(fmt.Sprintf("Statement '%s' not allowed in a pseudo function", stmtCtx.GetStart().GetText()), line)))
					}
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
			fmt.Println(r2d2Styles.WarningMessage(formatWarningMessage("Loop has no escape!", line)))
		}
	}

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitLoopStatement(ctx *parser.LoopStatementContext) any {
	v.JsCode += fmt.Sprintf("while (true){")

	result := v.VisitChildren(ctx)

	v.JsCode += "}"

	return result
}

func (v *R2D2Visitor) VisitFunctionCallStatement(ctx *parser.FunctionCallStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) any {

	// Skip if no identifier
	if ctx.IDENTIFIER() == nil {
		return v.VisitChildren(ctx)
	}
	varName := ctx.IDENTIFIER().GetText()
	// Create variable object
	variable := Variable{
		Name:       varName,
		isExported: ctx.EXPORT() != nil,
	}
	if ctx.TypeExpression() != nil {
		variable.Type = ctx.TypeExpression().GetText()
	}

	// Generate JS declaration
	if ctx.LET() != nil {
		v.JsCode += fmt.Sprintf("let %s", varName)
	} else if ctx.CONST() != nil {
		if ctx.ASSIGN() == nil {
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Const variable must be assigned a value")))
		}
		v.JsCode += fmt.Sprintf("const %s", varName)
	} else if ctx.VAR() != nil {
		v.JsCode += fmt.Sprintf("var %s", varName)
	}

	// Handle assignment if present
	if ctx.ASSIGN() != nil && ctx.Expression() != nil {
		v.JsCode += " = "
		ctx.Expression().Accept(v)
		variable.Value = "sei la" // Set the value before storing the variable
	}

	// Store the variable AFTER setting all its properties
	if _, ok := ctx.GetParent().(*parser.ModuleDeclarationContext); ok {
		v.currentModule.Variables[varName] = variable

		// add the ; if the variable is in module level
		defer func() {
			v.JsCode += ";"
		}()
	} else {
		// Check for invalid export
		if ctx.EXPORT() != nil {
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Variable cannot be exported in a function")))
		}
		v.currentFunction.Variables[varName] = variable
	}

	// v.JsCode += ";"
	return nil
}

func (v *R2D2Visitor) VisitStatement(ctx *parser.StatementContext) any {
	// Visit statement children first
	result := v.VisitChildren(ctx)

	// Add appropriate semicolons based on statement type
	switch ctx.GetChild(0).(type) {
	case *parser.ExpressionStatementContext,
		*parser.AssignmentDeclarationContext,
		*parser.FunctionCallStatementContext,
		*parser.CicleControlContext,
		*parser.ReturnStatementContext,
		*parser.VariableDeclarationContext:
		v.JsCode += ";"
	}

	return result
}

func (v *R2D2Visitor) VisitExpression(ctx *parser.ExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitBreakStatement(ctx *parser.BreakStatementContext) any {
	inLoop := findParent(ctx, (*parser.LoopStatementContext)(nil), (*parser.ForStatementContext)(nil), (*parser.WhileStatementContext)(nil))

	if !inLoop {
		errorMessage := fmt.Sprintf("Break statement on %s must be within a loop", r2d2Styles.Bold(fmt.Sprintf("line %d", ctx.GetStart().GetLine())))
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(errorMessage, ctx.GetStart().GetLine())))
	}

	v.JsCode += "break"
	return nil
}

func (v *R2D2Visitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) any {
	// Verify we're in a loop context
	inLoop := findParent(ctx, (*parser.LoopStatementContext)(nil), (*parser.ForStatementContext)(nil), (*parser.WhileStatementContext)(nil))

	if !inLoop {
		errorMessage := fmt.Sprintf("Continue statement on %s must be within a loop", r2d2Styles.Bold(fmt.Sprintf("line %d", ctx.GetStart().GetLine())))
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(errorMessage, ctx.GetStart().GetLine())))
	}

	v.JsCode += "continue"
	return nil
}

func (v *R2D2Visitor) VisitForStatement(ctx *parser.ForStatementContext) any {
	v.JsCode += "for ("

	simpleFor := ctx.SimpleFor()

	bothAssignments := simpleFor.Assignment(0) != nil && simpleFor.Assignment(1) != nil
	if simpleFor != nil {
		// Initialization
		if simpleFor.VariableDeclaration() != nil {
			simpleFor.VariableDeclaration().Accept(v)
		} else if bothAssignments {
			simpleFor.Assignment(0).Accept(v)
			// // Initial assignment
			// if simpleFor.Assignment(0) != nil {
			// 	assignment := simpleFor.Assignment(0)
			// 	v.JsCode += assignment.IDENTIFIER().GetText()

			// 	if assignment.AssignmentOperator() != nil {
			// 		v.JsCode += " " + assignment.AssignmentOperator().GetText() + " "
			// 		if assignment.Expression() != nil {
			// 			assignment.Expression().Accept(v)
			// 		}
			// 	} else if assignment.INCREMENT() != nil {
			// 		v.JsCode += "++"
			// 	} else if assignment.DECREMENT() != nil {
			// 		v.JsCode += "--"
			// 	}
			// }
		}

		// Condition
		v.JsCode += "; "
		if simpleFor.Expression() != nil {
			simpleFor.Expression().Accept(v)
		}

		// Update
		v.JsCode += "; "
		if simpleFor.AllAssignment() != nil {
			if bothAssignments {
				simpleFor.Assignment(1).Accept(v)
			} else {
				simpleFor.Assignment(0).Accept(v)
			}
		}
	}

	v.JsCode += ") {"

	// Process block
	if ctx.Block() != nil {
		ctx.Block().Accept(v)
	}

	v.JsCode += "}"

	return nil
}

func (v *R2D2Visitor) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) any {
	if ctx.IDENTIFIER() == nil {
		return v.VisitChildren(ctx)
	}

	typeName := ctx.IDENTIFIER().GetText()

	// Store the type declaration in the module's Types map
	v.currentModule.Types[typeName] = ctx

	// JavaScript doesn't have direct type declarations, so we'll just add a comment
	v.JsCode += fmt.Sprintf("/* Type declaration: %s */", typeName)

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) any {
	if ctx.IDENTIFIER() == nil {
		return v.VisitChildren(ctx)
	}

	interfaceName := ctx.IDENTIFIER().GetText()

	// Create a new interface
	newInterface := Interface{
		Name:      interfaceName,
		Functions: make(map[string]Function),
	}

	// Process function declarations in the interface
	for _, funcDecl := range ctx.AllFunctionDeclaration() {
		if funcDecl.IDENTIFIER() == nil {
			continue
		}

		funcName := funcDecl.IDENTIFIER().GetText()

		// Parse function arguments
		arguments := make(map[string]Argument)
		if funcDecl.ParameterList() != nil {
			for _, param := range funcDecl.ParameterList().AllParameter() {
				if param.IDENTIFIER() != nil && param.TypeExpression() != nil {
					paramName := param.IDENTIFIER().GetText()
					arguments[paramName] = Argument{
						Name: paramName,
						Type: param.TypeExpression().GetText(),
					}
				}
			}
		}

		// Store the function in the interface
		newInterface.Functions[funcName] = Function{
			Name:      funcName,
			Arguments: arguments,
			Variables: make(map[string]Variable),
			Functions: make(map[string]Function),
		}
	}

	// Add the interface to the symbol table
	v.symbolTable.Interfaces[interfaceName] = newInterface

	// JavaScript doesn't have interfaces, so add a comment
	v.JsCode += fmt.Sprintf("/* Interface declaration: %s */", interfaceName)

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) any {
	if ctx.Expression() != nil {
		ctx.Expression().Accept(v)
	}
	return nil
}

func (v *R2D2Visitor) VisitCicleControl(ctx *parser.CicleControlContext) any {

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) any {
	v.JsCode += "return"

	if ctx.Expression() != nil {
		v.JsCode += " "
		ctx.Expression().Accept(v)
	}

	return nil
}

// VisitIfStatement handles if statements
func (v *R2D2Visitor) VisitIfStatement(ctx *parser.IfStatementContext) any {
	v.JsCode += "if ("

	// Condição do primeiro if
	if ctx.Expression(0) != nil {
		expr := ctx.Expression(0).Accept(v)
		if exprText, ok := expr.(string); ok && exprText != "" {
			v.JsCode += exprText
		}
	}

	v.JsCode += ") "

	// Corpo do IF: bloco ou statement
	if ctx.Block(0) != nil {
		v.JsCode += "{"
		ctx.Block(0).Accept(v)
		v.JsCode += "}"
	} else if ctx.Statement(0) != nil {
		v.JsCode += "{"
		ctx.Statement(0).Accept(v)
		v.JsCode += "}"
	}

	// ELSE IFs
	elseIfCount := len(ctx.AllExpression()) - 1
	for i := 0; i < elseIfCount; i++ {
		v.JsCode += " else if ("
		expr := ctx.Expression(i + 1).Accept(v)
		if exprText, ok := expr.(string); ok && exprText != "" {
			v.JsCode += exprText
		}
		v.JsCode += ") "

		blockIndex := i + 1
		if len(ctx.AllBlock()) > blockIndex && ctx.Block(blockIndex) != nil {
			v.JsCode += "{"
			ctx.Block(blockIndex).Accept(v)
			v.JsCode += "}"
		} else if len(ctx.AllStatement()) > blockIndex && ctx.Statement(blockIndex) != nil {
			v.JsCode += "{"
			ctx.Statement(blockIndex).Accept(v)
			v.JsCode += "}"
		}
	}

	// ELSE final
	elseIndex := len(ctx.AllExpression())
	if len(ctx.AllBlock()) > elseIndex && ctx.ELSE(0) != nil {
		v.JsCode += " else "
		if ctx.Block(elseIndex) != nil {
			v.JsCode += "{"
			ctx.Block(elseIndex).Accept(v)
			v.JsCode += "}"
		} else if ctx.Statement(elseIndex) != nil {
			v.JsCode += "{"
			ctx.Statement(elseIndex).Accept(v)
			v.JsCode += "}"
		}
	}

	return nil
}

func (v *R2D2Visitor) VisitAssignmentDeclaration(ctx *parser.AssignmentDeclarationContext) any {
	// if ctx.Assignment() != nil {
	// 	return ctx.Assignment().Accept(v)
	// }
	return v.VisitChildren(ctx)
}

// VisitAssignment handles assignments
func (v *R2D2Visitor) VisitAssignment(ctx *parser.AssignmentContext) any {
	if ctx.IDENTIFIER() == nil {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage("Assignment identifier missing", ctx.GetStart().GetLine())))
		return v.VisitChildren(ctx)
	}

	// Process the assignment
	varName := ctx.IDENTIFIER().GetText()

	// Verify the variable exists in the current scope
	varExists := false
	var variabl Variable

	// Check in current function, module, and global variables
	for _, scope := range []map[string]Variable{
		v.currentFunction.Variables,
		v.currentModule.Variables,
		// v.symbolTable.Globals,
	} {
		if vari, exists := scope[varName]; exists {
			varExists = true
			variabl = vari
			break
		}
	}
	if !varExists {
		errorMessage := fmt.Sprintf("Variable '%s' not declared on line %s", r2d2Styles.Bold(varName), r2d2Styles.Bold(fmt.Sprintf("%d", ctx.GetStart().GetLine())))
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(errorMessage, ctx.GetStart().GetLine())))
		v.JsCode += fmt.Sprintf("/* ERROR: %s */", errorMessage)
		return nil
	}

	// Generate assignment code
	v.JsCode += varName

	// Handle assignment operator
	if ctx.AssignmentOperator() != nil {
		if ctx.AssignmentOperator().GetText() != "=" {
			if variabl.Value == nil || variabl.Value == "" {
				line := ctx.GetStart().GetLine()
				fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Variable '%s' not initialized on %s", ctx.IDENTIFIER().GetText(), r2d2Styles.Bold(fmt.Sprintf(" line %d", line)))))
			}
		}
		v.JsCode += " " + ctx.AssignmentOperator().GetText() + " "

		// Use the expression visitor for the right-hand side, without GetText
		if ctx.Expression() != nil {
			ctx.Expression().Accept(v)
		}
	} else if ctx.INCREMENT() != nil {
		// Handle increment operation
		v.JsCode += "++"
	} else if ctx.DECREMENT() != nil {
		// Handle decrement operation
		v.JsCode += "--"
	}

	return nil
}

// VisitWhileStatement handles while statements
func (v *R2D2Visitor) VisitWhileStatement(ctx *parser.WhileStatementContext) any {
	v.JsCode += "while ("

	// Condition using our improved expression visitor
	if ctx.Expression() != nil {
		exprResult := ctx.Expression().Accept(v)
		if exprText, ok := exprResult.(string); ok && exprText != "" {
			v.JsCode += exprText
		}
	}

	v.JsCode += ") {"

	// Body
	if ctx.Block() != nil {
		ctx.Block().Accept(v)
	}

	v.JsCode += "}"

	return nil
}

// VisitSwitchStatement handles switch statements
func (v *R2D2Visitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) any {
	v.JsCode += "switch ("

	// Expressão usando nosso visitante de expressão aprimorado
	if ctx.Expression() != nil {
		exprResult := ctx.Expression().Accept(v)
		if exprText, ok := exprResult.(string); ok && exprText != "" {
			v.JsCode += exprText
		}
	}

	v.JsCode += ") {"

	// Processar os casos
	for _, switchCase := range ctx.AllSwitchCase() {
		v.JsCode += "case "
		if switchCase.Expression() != nil {
			exprResult := switchCase.Expression().Accept(v)
			if exprText, ok := exprResult.(string); ok && exprText != "" {
				v.JsCode += exprText
			}
		}
		v.JsCode += ":"
		if switchCase.Block() != nil {
			switchCase.Block().Accept(v)
		} else if switchCase.Statement() != nil {
			switchCase.Statement().Accept(v)
		}
		v.JsCode += "break;"
	}

	// Caso default
	if ctx.DefaultCase() != nil {
		v.JsCode += "default:"
		if ctx.DefaultCase().Block() != nil {
			ctx.DefaultCase().Block().Accept(v)
		} else if ctx.DefaultCase().Statement() != nil {
			ctx.DefaultCase().Statement().Accept(v)
		}
	}

	v.JsCode += "}"

	return nil
}

// VisitFunctionCall handles function calls
func (v *R2D2Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) any {
	// Get function name (handle both simple and qualified names)
	var funcName string
	if ctx.IDENTIFIER(0) == nil {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Missing function identifier")))
		v.JsCode += "/* ERROR: Missing function identifier */"
		return nil
	}

	funcName = ctx.IDENTIFIER(0).GetText()

	// Check for qualified name (e.g., console.log)
	if len(ctx.AllIDENTIFIER()) > 1 && ctx.IDENTIFIER(1) != nil {
		funcName = funcName + "." + ctx.IDENTIFIER(1).GetText()
	}

	isAccessible, fn := isAccessibleFunction(v, funcName)
	if !isAccessible {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(fmt.Sprintf("Function '%s' not found", funcName), ctx.GetStart().GetLine())))
		return nil
	}

	// Generate JS code for function call
	if strings.Contains(funcName, ".") {
		v.JsCode += funcName
	} else if v.currentModule.Name != "" && v.currentModule.Name != strings.Split(funcName, ".")[0] {
		v.JsCode += funcName
	} else {
		v.JsCode += funcName
	}

	v.JsCode += "("

	// Visit arguments properly
	argumentList := ctx.ArgumentList()

	if argumentList != nil {
		exprs := argumentList.AllExpression()
		if len(exprs) != len(fn.Arguments) {
			lineNum := ctx.GetStart().GetLine()
			var msg string
			switch len(exprs) > len(fn.Arguments) {
			case true:
				msg = fmt.Sprintf(
					"Too many arguments passed when calling function %s at line %s",
					r2d2Styles.Bold(funcName),
					r2d2Styles.Bold(fmt.Sprintf("%d", lineNum)),
				)

			case false:

				msg = fmt.Sprintf(
					"Too little arguments passed when calling function %s at line %s",
					r2d2Styles.Bold(funcName),
					r2d2Styles.Bold(fmt.Sprintf("%d", lineNum)),
				)
			}
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, ctx.GetStart().GetLine())))
		}
		for i, expr := range exprs {
			expr.Accept(v) // visita normalmente o argumento
			if i < len(exprs)-1 {
				v.JsCode += ", "
			}
		}
	}

	v.JsCode += ")"

	return nil
}

func (v *R2D2Visitor) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) any {
	if ctx.Literal() != nil {
		v.JsCode += ctx.Literal().GetText()
		// fmt.Println(r2d2Styles.InfoMessage("Literal found: " + r2d2Styles.Bold(ctx.Literal().GetText())))
	} else {
		// fmt.Println(r2d2Styles.ErrorMessage("Literal not found"))
	}
	return nil
}

func (v *R2D2Visitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) any {
	line := ctx.GetStart().GetLine()
	if ctx.IDENTIFIER() != nil {
		if ok, variable := isAccessibleVariable(v, ctx.IDENTIFIER().GetText()); ok {
			if variable.Value == nil || variable.Value == "" {
				fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Variable '%s' not initialized on %s", ctx.IDENTIFIER().GetText(), r2d2Styles.Bold(fmt.Sprintf(" line %d", line)))))
			}
			v.JsCode += ctx.IDENTIFIER().GetText()
		} else {
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(fmt.Sprintf("Variable '%s' not found", ctx.IDENTIFIER().GetText()), line)))
		}
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage("Identifier not found", line)))
	}
	return nil
}

func (v *R2D2Visitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) any {
	ctx.FunctionCall().Accept(v)
	return nil
}

func (v *R2D2Visitor) VisitParenthesisExpression(ctx *parser.ParenthesisExpressionContext) any {
	v.JsCode += "("
	ctx.Expression().Accept(v)
	v.JsCode += ")"
	return nil
}

func (v *R2D2Visitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) any {
	if opNode, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		token := opNode.GetSymbol().GetTokenType()

		switch token {
		case parser.R2D2ParserNOT:
			v.JsCode += "!"
		case parser.R2D2ParserMINUS:
			v.JsCode += "-"
		case parser.R2D2ParserINCREMENT:
			v.JsCode += "++"
		case parser.R2D2ParserDECREMENT:
			v.JsCode += "--"
		default:
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Unknown unary operator")))
		}
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Expected a unary operator")))
	}

	if ctx.Expression() != nil {
		ctx.Expression().Accept(v)
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Expression not found")))
	}

	return nil
}

func (v *R2D2Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) any {
	if ctx.Expression(0) != nil {
		ctx.Expression(0).Accept(v)
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage("Left expression not found", ctx.GetStart().GetLine())))
	}

	// Operator is child(1)
	if opNode, ok := ctx.GetChild(1).(antlr.TerminalNode); ok {
		token := opNode.GetSymbol().GetTokenType()

		switch token {
		case parser.R2D2ParserMULT:
			v.JsCode += "*"
		case parser.R2D2ParserDIV:
			v.JsCode += "/"
		case parser.R2D2ParserMOD:
			v.JsCode += "%"
		default:
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Unknown multiplicative operator")))
		}
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Expected a multiplicative operator")))
	}

	if ctx.Expression(1) != nil {
		ctx.Expression(1).Accept(v)
	} else {
		fmt.Println(r2d2Styles.ErrorMessage("Right expression not found"))
	}

	return nil
}

func (v *R2D2Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) any {
	if ctx.Expression(0) != nil {
		ctx.Expression(0).Accept(v)
	} else {
		fmt.Println(r2d2Styles.ErrorMessage("Left expression not found"))
	}

	// Operator is child(1)
	if opNode, ok := ctx.GetChild(1).(antlr.TerminalNode); ok {
		token := opNode.GetSymbol().GetTokenType()

		switch token {
		case parser.R2D2ParserPLUS:
			v.JsCode += "+"
		case parser.R2D2ParserMINUS:
			v.JsCode += "-"
		default:
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage("Unknown additive operator", ctx.GetStart().GetLine())))
		}
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage("Expected an additive operator", ctx.GetStart().GetLine())))
	}

	if ctx.Expression(1) != nil {
		ctx.Expression(1).Accept(v)
	} else {
		fmt.Println(r2d2Styles.ErrorMessage("Right expression not found"))
	}

	return nil
}

func (v *R2D2Visitor) VisitComparisonExpression(ctx *parser.ComparisonExpressionContext) any {
	if ctx.Expression(0) != nil {
		ctx.Expression(0).Accept(v)
	} else {
		fmt.Println(r2d2Styles.ErrorMessage("Left expression not found"))
	}

	if opNode, ok := ctx.GetChild(1).(antlr.TerminalNode); ok {
		token := opNode.GetSymbol().GetTokenType()

		switch token {
		case parser.R2D2ParserEQ:
			v.JsCode += "=="
		case parser.R2D2ParserNEQ:
			v.JsCode += "!="
		case parser.R2D2ParserLT:
			v.JsCode += "<"
		case parser.R2D2ParserGT:
			v.JsCode += ">"
		case parser.R2D2ParserLEQ:
			v.JsCode += "<="
		case parser.R2D2ParserGEQ:
			v.JsCode += ">="
		default:
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Unknown comparison operator")))
		}
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Expected a comparison operator")))
	}

	if ctx.Expression(1) != nil {
		ctx.Expression(1).Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage("Right expression not found", line)))
	}

	return nil
}

func (v *R2D2Visitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) any {
	if ctx.Expression(0) != nil {
		ctx.Expression(0).Accept(v)
	} else {
		fmt.Println(r2d2Styles.ErrorMessage("Left expression not found"))
	}

	if opNode, ok := ctx.GetChild(1).(antlr.TerminalNode); ok {
		token := opNode.GetSymbol().GetTokenType()

		switch token {
		case parser.R2D2ParserAND:
			v.JsCode += "&&"
		case parser.R2D2ParserOR:
			v.JsCode += "||"
		default:
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Unknown logical operator")))
		}
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Expected a logical operator")))
	}

	if ctx.Expression(1) != nil {
		ctx.Expression(1).Accept(v)
	} else {
		fmt.Println(r2d2Styles.ErrorMessage("Right expression not found"))
	}

	return nil
}

func (v *R2D2Visitor) VisitJsStatement(ctx *parser.JsStatementContext) any {
	if ctx.JS_BLOCK() != nil {
		raw := ctx.JS_BLOCK().GetText()
		// Remove os delimitadores '<<' e '>>'
		clean := strings.TrimSuffix(strings.TrimPrefix(raw, "<<"), ">>")
		v.JsCode += clean
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage("JS code not found", ctx.GetStart().GetLine())))
	}

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitArrayAccessExpression(ctx *parser.ArrayAccessExpressionContext) any {
	if ctx.Expression(0) != nil {
		ctx.Expression(0).Accept(v)
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine("Array not found")))
	}

	if ctx.LBRACK() != nil {
		v.JsCode += "["
	}

	if ctx.Expression(1) != nil {
		ctx.Expression(1).Accept(v)
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage("Index not found", ctx.GetStart().GetLine())))
	}

	if ctx.RBRACK() != nil {
		v.JsCode += "]"
	}

	return nil
}

func (v *R2D2Visitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) any {
	if ctx.LBRACK() != nil {
		v.JsCode += "["
	}

	// Use AllExpression() to get the slice of expressions
	for i, expr := range ctx.AllExpression() {
		if i > 0 {
			v.JsCode += ", "
		}
		expr.Accept(v)
	}

	if ctx.RBRACK() != nil {
		v.JsCode += "]"
	}

	return nil
}
