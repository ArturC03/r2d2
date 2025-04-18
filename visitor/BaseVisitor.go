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
	Arguments  []Argument
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
func isExported(node interface{}) bool {
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

func findChild(parent antlr.RuleContext, types ...interface{}) bool {
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

func findParent(node antlr.RuleContext, types ...interface{}) bool {
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

// func loadGlobalFunctions(v *R2D2Visitor) error {
// 	// Add standard library functions to the symbol table if they don't exist
// // 	standardFunctions := map[string][]Argument{
// 		"console.log": {
// 			{Name: "message", Type: "any"},
// 		},
// 		"console.error": {
// 			{Name: "message", Type: "any"},
// 		},
// 		"console.warn": {
// 			{Name: "message", Type: "any"},
// 		},
// 		"Math.random": {},
// 		"Math.floor": {
// 			{Name: "value", Type: "number"},
// 		},
// 		// Add more standard functions as needed
// 	}

// Create console module if it doesn't exist
// 	if _, exists := v.symbolTable.Modules["console"]; !exists {
// 		v.symbolTable.Modules["console"] = Module{
// 			Name:      "console",
// 			Functions: make(map[string]Function),
// 			Variables: make(map[string]Variable),
// 			Types:     make(map[string]any),
// 		}
// 	}

// 	// Create Math module if it doesn't exist
// 	if _, exists := v.symbolTable.Modules["Math"]; !exists {
// 		v.symbolTable.Modules["Math"] = Module{
// 			Name:      "Math",
// 			Functions: make(map[string]Function),
// 			Variables: make(map[string]Variable),
// 			Types:     make(map[string]any),
// 		}
// 	}

// 	// Add functions to their respective modules
// 	for fullName, args := range standardFunctions {
// 		parts := strings.Split(fullName, ".")
// 		if len(parts) == 2 {
// 			moduleName, funcName := parts[0], parts[1]

// 			if module, exists := v.symbolTable.Modules[moduleName]; exists {
// 				module.Functions[funcName] = Function{
// 					Name:       funcName,
// 					Arguments:  args,
// 					Variables:  make(map[string]Variable),
// 					Functions:  make(map[string]Function),
// 					isExported: true,
// 				}
// 				v.symbolTable.Modules[moduleName] = module
// 			}
// 		}
// 	}

// 	return nil
// }

// func (v *R2D2Visitor) isAccessibleFunction(funcName string) (bool, Function, string) {
// 	// Check if function is a method call (contains a dot)
// 	parts := strings.Split(funcName, ".")

// 	if len(parts) == 2 {
// 		// Module.function format
// 		moduleName, methodName := parts[0], parts[1]

// 		// Check if module exists
// 		module, moduleExists := v.symbolTable.Modules[moduleName]
// 		if !moduleExists {
// 			return false, Function{}, fmt.Sprintf("/* ERROR: Module '%s' not found */", moduleName)
// 		}

// 		// Check if function exists in module
// 		function, functionExists := module.Functions[methodName]
// 		if !functionExists {
// 			return false, Function{}, fmt.Sprintf("/* ERROR: Function '%s' not found in module '%s' */", methodName, moduleName)
// 		}

// 		// Check if function is exported (if calling from another module)
// 		if moduleName != v.currentModule.Name && !function.isExported {
// 			return false, Function{}, fmt.Sprintf("/* ERROR: Function '%s' is not exported from module '%s' */", methodName, moduleName)
// 		}

// 		return true, function, ""
// 	} else {
// 		// Check current module functions
// 		if function, exists := v.currentModule.Functions[funcName]; exists {
// 			return true, function, ""
// 		}

// 		// Check global functions
// 		if _, exists := v.symbolTable.Globals[funcName]; exists {
// 			// Convert Global to Function for interface compatibility
// 			return true, Function{
// 				Name:       funcName,
// 				Arguments:  []Argument{},
// 				isExported: true,
// 			}, ""
// 		}

// 		// Function not found
// 		return false, Function{}, fmt.Sprintf("/* ERROR: Function '%s' not found */", funcName)
// 	}
// }

// func isValidJSType(value string, expectedType string) bool {
// 	// Basic type checking - this could be expanded with more sophisticated checks
// 	if expectedType == "any" {
// 		return true
// 	}

// 	// Check for string literals
// 	if expectedType == "string" && (strings.HasPrefix(value, "\"") || strings.HasPrefix(value, "'")) {
// 		return true
// 	}

// 	// Check for number literals
// 	if expectedType == "number" {
// 		// Simple check for numeric format
// 		if _, err := fmt.Sscanf(value, "%f", new(float64)); err == nil {
// 			return true
// 		}
// 	}

// 	// Check for boolean literals
// 	if expectedType == "boolean" && (value == "true" || value == "false") {
// 		return true
// 	}

// 	// For objects, arrays, and other complex types, more sophisticated checks would be needed
// 	// Assuming valid for now (can be enhanced later)
// 	return true
// }

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
	// Check if any of the required context objects are nil
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

	// Store global in the symbol table with proper information
	globalName := ctx.IDENTIFIER().GetText()
	v.symbolTable.Globals[globalName] = Global{
		Name:  globalName,
		Value: ctx.Expression().GetText(),
		Type:  ctx.TypeExpression().GetText(),
	}

	jsCode := fmt.Sprintf("const %s = %s;", globalName, ctx.Expression().GetText())
	v.JsCode += jsCode

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitModuleDeclaration(ctx *parser.ModuleDeclarationContext) any {
	// Make sure this is in your initialization code for R2D2Visitor
	if v.symbolTable.Globals == nil {
		v.symbolTable.Globals = make(map[string]Global)
	}

	// Add null check for IDENTIFIER
	if ctx.IDENTIFIER(0) == nil || len(ctx.AllIDENTIFIER()) == 0 {
		fmt.Println(r2d2Styles.ErrorMessage("Error: Module declaration without identifier"))
		return nil
	}

	moduleName := ctx.IDENTIFIER(0).GetText()

	// Create Module
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
		fmt.Println(r2d2Styles.ErrorMessage("Module " + moduleName + " already exists"))
		return nil
	}

	// Process module contents to populate the symbol table
	for _, child := range ctx.GetChildren() {
		// Function Declaration
		if funcDecl, ok := child.(*parser.FunctionDeclarationContext); ok {
			if funcDecl.IDENTIFIER() == nil {
				continue
			}

			funcName := funcDecl.IDENTIFIER().GetText()

			// Parse function arguments
			var arguments []Argument
			if funcDecl.ParameterList() != nil {
				for _, param := range funcDecl.ParameterList().AllParameter() {
					if param.IDENTIFIER() != nil && param.TypeExpression() != nil {
						arguments = append(arguments, Argument{
							Name: param.IDENTIFIER().GetText(),
							Type: param.TypeExpression().GetText(),
						})
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

			// Store function in the module
			module := v.symbolTable.Modules[moduleName]
			module.Functions[funcName] = function
			v.symbolTable.Modules[moduleName] = module
		}

		// Variable Declaration
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

			// Variable Declaration with Assignment
			if varDecl.ASSIGN() != nil && varDecl.Expression() != nil {
				variable.Value = varDecl.Expression().GetText()
			} else {
				variable.Value = ""
			}

			// Store variable in the module
			module := v.symbolTable.Modules[moduleName]
			module.Variables[varName] = variable
			v.symbolTable.Modules[moduleName] = module
		}

		// Type Declaration
		if typeDecl, ok := child.(*parser.TypeDeclarationContext); ok {
			if typeDecl.IDENTIFIER() == nil {
				continue
			}

			typeName := typeDecl.IDENTIFIER().GetText()

			// Store type in the module
			module := v.symbolTable.Modules[moduleName]
			module.Types[typeName] = typeDecl
			v.symbolTable.Modules[moduleName] = module
		}
	}

	// Start Module in JS code
	v.JsCode += fmt.Sprintf("const %s = (function () {\n", moduleName)

	// Visit children to process the module contents
	result := v.VisitChildren(ctx)

	// Get exported items for module return statement
	moduleExports := v.symbolTable.Modules[moduleName].Exports()

	// End Module with return statement
	v.JsCode += fmt.Sprintf("return {%s}; })();\n", strings.Join(moduleExports, ", "))

	// Check if this module exports a main function and execute it
	module := v.symbolTable.Modules[moduleName]
	if mainFunc, exists := module.Functions["main"]; exists && mainFunc.isExported {
		v.JsCode += fmt.Sprintf("// Auto-execute exported main function\n%s.main();\n", moduleName)
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
		Arguments:  []Argument{},
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
				paramNames = append(paramNames, param.IDENTIFIER().GetText())
				variable := Variable{
					Name: param.IDENTIFIER().GetText(),
					Type: param.TypeExpression().GetText(),
				}
				v.currentFunction.Variables[variable.Name] = variable
				fmt.Println(v.currentFunction.Variables)
			}
		}

		v.JsCode += strings.Join(paramNames, ", ")
	}

	v.JsCode += ") {\n"

	// Store any local variables defined in the function's block
	if ctx.Block() != nil {
		// Visit the function body
		ctx.Block().Accept(v)
	}

	v.JsCode += "}\n"

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
						fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Line %s: statement %s not allowed in a pseudo function", r2d2Styles.Bold(fmt.Sprintf("%d", line)), r2d2Styles.Bold(stmtCtx.GetStart().GetText()))))
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
}

func (v *R2D2Visitor) VisitFunctionCallStatement(ctx *parser.FunctionCallStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) any {
	// Load global functions if needed
	err := loadGlobalFunctions(v)
	if err != nil {
		v.JsCode += fmt.Sprintf("/* ERROR: %s */", err.Error())
		return nil
	}

	// Get function name (handle both simple and qualified names)
	var funcName string
	if ctx.IDENTIFIER(0) == nil {
		v.JsCode += "/* ERROR: Missing function identifier */"
		return nil
	}

	funcName = ctx.IDENTIFIER(0).GetText()

	// Check for qualified name (e.g., console.log)
	if len(ctx.AllIDENTIFIER()) > 1 && ctx.IDENTIFIER(1) != nil {
		funcName = funcName + "." + ctx.IDENTIFIER(1).GetText()
	}

	// Verify function accessibility
	isAccessible, function, errorMessage := v.isAccessibleFunction(funcName)
	if !isAccessible {
		v.JsCode += errorMessage + "\n"
		fmt.Println(r2d2Styles.ErrorMessage(errorMessage))
		return nil
	}

	// Get arguments
	var passedArgs []string
	argumentList := ctx.ArgumentList()
	if argumentList != nil {
		for _, arg := range argumentList.AllExpression() {
			passedArgs = append(passedArgs, arg.GetText())
		}
	}

	// Verify argument count for non-variadic functions
	if len(passedArgs) != len(function.Arguments) {
		// List of functions that accept variable number of arguments
		variableArgsAllowed := []string{
			"console.log", "console.error", "console.warn", "console.info",
			"Array.push", "Array.concat",
			"Math.max", "Math.min",
			"Object.assign",
			"setTimeout", "setInterval",
			"Function.apply", "Function.call",
		}

		// Check if this function allows variable args
		isVariableArgsFunction := false
		for _, varArgFunc := range variableArgsAllowed {
			if funcName == varArgFunc || strings.HasSuffix(funcName, ".apply") || strings.HasSuffix(funcName, ".call") {
				isVariableArgsFunction = true
				break
			}
		}

		// Show error if argument count mismatch for non-variadic functions
		if !isVariableArgsFunction {
			errorMessage := fmt.Sprintf(
				"/* ERROR: Function '%s' expects %d arguments, but %d were provided */",
				funcName, len(function.Arguments), len(passedArgs),
			)
			v.JsCode += errorMessage + "\n"
			fmt.Println(r2d2Styles.ErrorMessage(errorMessage))
			return nil
		}
	}

	// Type checking for arguments
	for i, passedArg := range passedArgs {
		if i < len(function.Arguments) {
			expectedType := function.Arguments[i].Type
			if !isValidJSType(passedArg, expectedType) {
				errorMessage := fmt.Sprintf(
					"/* ERROR: Argument %d of function '%s' expects type '%s', but got '%s' */",
					i+1, funcName, expectedType, passedArg,
				)
				v.JsCode += errorMessage + "\n"
				fmt.Println(r2d2Styles.ErrorMessage(errorMessage))
				return nil
			}
		}
	}

	// Generate JS code for function call
	// Handle qualified vs unqualified names differently
	if strings.Contains(funcName, ".") {
		v.JsCode += fmt.Sprintf("%s(%s)", funcName, strings.Join(passedArgs, ", "))
	} else if v.currentModule.Name != strings.Split(funcName, ".")[0] {
		v.JsCode += fmt.Sprintf("%s.%s(%s)", v.currentModule.Name, funcName, strings.Join(passedArgs, ", "))
	} else {
		v.JsCode += fmt.Sprintf("%s(%s)", funcName, strings.Join(passedArgs, ", "))
	}

	return nil
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

	// Handle assignment if present
	if ctx.ASSIGN() != nil && ctx.Expression() != nil {
		variable.Value = ctx.Expression().GetText()
	}

	// If this is a module-level variable
	if _, ok := ctx.GetParent().(*parser.ModuleDeclarationContext); ok {
		// Check for invalid export
		if ctx.EXPORT() != nil {
			fmt.Println(r2d2Styles.ErrorMessage("Cannot export non-global variables"))
		}
		// Store in module scope
		v.currentModule.Variables[varName] = variable
	} else {
		// Store in function scope
		v.currentFunction.Variables[varName] = variable
	}

	// Generate JS declaration based on type
	if ctx.LET() != nil {
		v.JsCode += fmt.Sprintf("let %s", varName)
	} else if ctx.CONST() != nil {
		if ctx.ASSIGN() == nil {
			fmt.Println(r2d2Styles.ErrorMessage("Const variable must be assigned a value"))
		} else {
			v.JsCode += fmt.Sprintf("const %s", varName)
		}
	} else if ctx.VAR() != nil {
		v.JsCode += fmt.Sprintf("var %s", varName)
	}

	// Handle assignment if present
	if ctx.ASSIGN() != nil && ctx.Expression() != nil {
		v.JsCode += fmt.Sprintf(" = %s", ctx.Expression().GetText())
	}

	return v.VisitChildren(ctx)
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

func (v *R2D2Visitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) any {
	v.JsCode += "return"

	// Add the return expression if present
	if ctx.Expression() != nil {
		// Get the expression result properly
		exprResult := ctx.Expression().Accept(v)

		// If the expression visitor returned a string, use it
		if exprText, ok := exprResult.(string); ok && exprText != "" {
			v.JsCode += " " + exprText
		} else {
			// Otherwise, get the text directly from the expression context
			v.JsCode += " " + ctx.Expression().GetText()
		}
	}

	return nil
}

func (v *R2D2Visitor) VisitIfStatement(ctx *parser.IfStatementContext) any {
	v.JsCode += "if ("

	// Expression for the condition
	if ctx.Expression(0) != nil {
		v.JsCode += ctx.Expression(0).GetText()
	}

	v.JsCode += ") {\n"

	// If block
	if ctx.Block(0) != nil {
		ctx.Block(0).Accept(v)
	}

	v.JsCode += "}"

	// Else block if present
	if len(ctx.AllBlock()) > 1 && ctx.ELSE(0) != nil {
		v.JsCode += " else {\n"
		ctx.Block(1).Accept(v)
		v.JsCode += "}"
	}

	return nil
}

func (v *R2D2Visitor) VisitExpression(ctx *parser.ExpressionContext) any {
	// For expressions, we'll just use the raw text for now
	// In a real compiler, you would build an expression tree
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitAssignmentDeclaration(ctx *parser.AssignmentDeclarationContext) any {
	if ctx.Assignment() != nil {
		// Process the assignment
		assignment := ctx.Assignment()
		if assignment.IDENTIFIER() == nil {
			return v.VisitChildren(ctx)
		}

		varName := assignment.IDENTIFIER().GetText()

		// Verify the variable exists in current scope
		varExists := false

		// First check in current function's variables (local scope)
		if _, exists := v.currentFunction.Variables[varName]; exists {
			varExists = true
		} else if _, exists := v.currentModule.Variables[varName]; exists {
			// Then check module variables (module scope)
			varExists = true
		} else if _, exists := v.symbolTable.Globals[varName]; exists {
			// Finally check global variables
			varExists = true
		}

		if !varExists {
			errorMessage := fmt.Sprintf("Variable '%s' not declared", varName)
			fmt.Println(r2d2Styles.ErrorMessage(errorMessage))
			v.JsCode += fmt.Sprintf("/* ERROR: %s */", errorMessage)
			return nil
		}

		// Rest of the code remains the same...
	}

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitBreakStatement(ctx *parser.BreakStatementContext) any {
	// Verify we're in a loop context
	v.JsCode += "break"
	return nil
}

func (v *R2D2Visitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) any {
	// Verify we're in a loop context
	inLoop := findParent(ctx, (*parser.LoopStatementContext)(nil), (*parser.ForStatementContext)(nil), (*parser.WhileStatementContext)(nil))

	if !inLoop {
		errorMessage := fmt.Sprintf("Continue statement on line %d must be within a loop", ctx.GetStart().GetLine())
		fmt.Println(r2d2Styles.ErrorMessage(errorMessage))
	}

	v.JsCode += "continue"
	return nil
}

func (v *R2D2Visitor) VisitForStatement(ctx *parser.ForStatementContext) any {
	v.JsCode += "for ("

	simpleFor := ctx.SimpleFor()
	if simpleFor != nil {
		// Initialization
		if simpleFor.VariableDeclaration() != nil {
			varDecl := simpleFor.VariableDeclaration()

			// Variable declaration type
			if varDecl.VAR() != nil {
				v.JsCode += "var "
			} else if varDecl.LET() != nil {
				v.JsCode += "let "
			} else if varDecl.CONST() != nil {
				v.JsCode += "const "
			}

			// Identifier
			v.JsCode += varDecl.IDENTIFIER().GetText()

			// Assignment
			if varDecl.ASSIGN() != nil && varDecl.Expression() != nil {
				v.JsCode += " = " + varDecl.Expression().GetText()
			}
		} else if simpleFor.Assignment(0) != nil {
			// Initial assignment
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

		// Condition
		v.JsCode += "; "
		if simpleFor.Expression() != nil {
			v.JsCode += simpleFor.Expression().GetText()
		}

		// Update
		v.JsCode += "; "
		if simpleFor.AllAssignment() != nil && len(simpleFor.AllAssignment()) > 0 {
			assignment := simpleFor.Assignment(len(simpleFor.AllAssignment()) - 1)
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

func (v *R2D2Visitor) VisitWhileStatement(ctx *parser.WhileStatementContext) any {
	v.JsCode += "while ("

	// Condition
	if ctx.Expression() != nil {
		v.JsCode += ctx.Expression().GetText()
	}

	v.JsCode += ") {\n"

	// Body
	if ctx.Block() != nil {
		ctx.Block().Accept(v)
	}

	v.JsCode += "}\n"

	return nil
}

func (v *R2D2Visitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) any {
	v.JsCode += "switch ("

	// Expression
	if ctx.Expression() != nil {
		v.JsCode += ctx.Expression().GetText()
	}

	v.JsCode += ") {\n"

	// Cases
	for _, switchCase := range ctx.AllSwitchCase() {
		v.JsCode += "case "
		if switchCase.Expression() != nil {
			v.JsCode += switchCase.Expression().GetText()
		}
		v.JsCode += ":\n"
		if switchCase.Block() != nil {
			switchCase.Block().Accept(v)
		}
		v.JsCode += "break;\n"
	}

	// Default case
	if ctx.DefaultCase() != nil {
		v.JsCode += "default:\n"
		if ctx.DefaultCase().Block() != nil {
			ctx.DefaultCase().Block().Accept(v)
		}
	}

	v.JsCode += "}\n"

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
	v.JsCode += fmt.Sprintf("/* Type declaration: %s */\n", typeName)

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
		var arguments []Argument
		if funcDecl.ParameterList() != nil {
			for _, param := range funcDecl.ParameterList().AllParameter() {
				if param.IDENTIFIER() != nil && param.TypeExpression() != nil {
					arguments = append(arguments, Argument{
						Name: param.IDENTIFIER().GetText(),
						Type: param.TypeExpression().GetText(),
					})
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
	v.JsCode += fmt.Sprintf("/* Interface declaration: %s */\n", interfaceName)

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) any {
	if ctx.Expression() != nil {
		v.JsCode += ctx.Expression().GetText()
	}
	return nil
}

func (v *R2D2Visitor) VisitCicleControl(ctx *parser.CicleControlContext) any {
	inLoop := findParent(ctx, (*parser.LoopStatementContext)(nil), (*parser.ForStatementContext)(nil), (*parser.WhileStatementContext)(nil))

	if !inLoop {
		errorMessage := fmt.Sprintf("Cicle control statement on line %d must be within a loop", ctx.GetStart().GetLine())
		fmt.Println(r2d2Styles.ErrorMessage(errorMessage))
	}

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) any {
	if ctx.ComparisonExpression(0) != nil {
		leftResult := v.Visit(ctx.ComparisonExpression(0))
		leftText, ok := leftResult.(string)
		if ok {
			v.JsCode += leftText
			for i := 0; i < ctx.GetChildCount(); i++ {
				opNode := ctx.GetChild(i)
				op, ok := opNode.(antlr.TerminalNode)
				if ok && (op.GetSymbol().GetTokenType() == parser.R2D2ParserAND || op.GetSymbol().GetTokenType() == parser.R2D2ParserOR) {
					rightIndex := i/2 + 1
					if ctx.ComparisonExpression(rightIndex) != nil {
						rightResult := v.Visit(ctx.ComparisonExpression(rightIndex))
						rightText, ok := rightResult.(string)
						if ok {
							v.JsCode += " (" + leftText + " " + op.GetText() + " " + rightText + ")"
							leftText = "(" + leftText + " " + op.GetText() + " " + rightText + ")" // Update for next iteration
						}
					}
				}
			}
		}
	}
	return nil // We are modifying v.JsCode directly
}

func (v *R2D2Visitor) VisitComparisonExpression(ctx *parser.ComparisonExpressionContext) any {
	if ctx.AdditiveExpression(0) != nil {
		leftResult := v.Visit(ctx.AdditiveExpression(0))
		leftText, ok := leftResult.(string)
		if ok {
			v.JsCode += leftText
			for i := 1; i < len(ctx.AllAdditiveExpression()); i++ {
				opNode := ctx.GetChild(2*i - 1)
				op, ok := opNode.(antlr.TerminalNode)
				if ok {
					rightResult := v.Visit(ctx.AdditiveExpression(i))
					rightText, ok := rightResult.(string)
					if ok {
						v.JsCode += " " + op.GetText() + " " + rightText
						leftText += " " + op.GetText() + " " + rightText // Update for potential parent expressions
					}
				}
			}
		}
	}
	return nil
}

func (v *R2D2Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) any {
	// Don't add to JsCode here, just build and return the expression text
	if ctx.MultiplicativeExpression(0) != nil {
		leftResult := v.Visit(ctx.MultiplicativeExpression(0))
		leftText, ok := leftResult.(string)
		if !ok {
			return ""
		}

		result := leftText
		for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
			opNode := ctx.GetChild(2*i - 1)
			op, ok := opNode.(antlr.TerminalNode)
			if !ok {
				continue
			}

			rightResult := v.Visit(ctx.MultiplicativeExpression(i))
			rightText, ok := rightResult.(string)
			if !ok {
				continue
			}

			result += " " + op.GetText() + " " + rightText
		}

		return result
	}
	return ""
}

func (v *R2D2Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) any {
	// Similar approach - don't add to JsCode, just return the text
	if ctx.UnaryExpression(0) != nil {
		leftResult := v.Visit(ctx.UnaryExpression(0))
		leftText, ok := leftResult.(string)
		if !ok {
			return ""
		}

		result := leftText
		for i := 1; i < len(ctx.AllUnaryExpression()); i++ {
			opNode := ctx.GetChild(2*i - 1)
			op, ok := opNode.(antlr.TerminalNode)
			if !ok {
				continue
			}

			rightResult := v.Visit(ctx.UnaryExpression(i))
			rightText, ok := rightResult.(string)
			if !ok {
				continue
			}

			result += " " + op.GetText() + " " + rightText
		}

		return result
	}
	return ""
}

func (v *R2D2Visitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) any {
	if ctx.GetChildCount() == 2 {
		opNode := ctx.GetChild(0)
		op, ok := opNode.(antlr.TerminalNode)
		if ok {
			exprResult := v.Visit(ctx.UnaryExpression())
			exprText, ok := exprResult.(string)
			if ok {
				return op.GetText() + exprText
			}
		}
		return ""
	}
	return v.Visit(ctx.MemberExpression())
}

func (v *R2D2Visitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) any {
	switch {
	case ctx.IDENTIFIER() != nil:
		return ctx.IDENTIFIER().GetText()

	case ctx.LPAREN() != nil:
		result := v.Visit(ctx.Expression())
		text, ok := result.(string)
		if ok {
			return "(" + text + ")"
		}
		return ""

	case ctx.Literal() != nil:
		return ctx.Literal().GetText()

	case ctx.ArrayLiteral() != nil:
		return v.Visit(ctx.ArrayLiteral())

	case ctx.FunctionCall() != nil:
		// Special case - FunctionCall already adds to JsCode
		v.Visit(ctx.FunctionCall())
		// Get the most recent part of JsCode that was added by the function call
		// This is a simplification - you might need a more robust method
		return "" // Return empty since JsCode was already modified

	default:
		return ""
	}
}
