package visitor

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"strings"

	"github.com/ArturC03/r2d2/parser"
	"github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
)

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
	// v.RegisterInterfaceNames(ctx)

	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitDeclaration(ctx *parser.DeclarationContext) any {
	return v.VisitChildren(ctx)
}

func (v *R2D2Visitor) VisitImportDeclaration(ctx *parser.ImportDeclarationContext) any {

	// Check if the file path is present
	if ctx.STRING_LITERAL() == nil {
		msg := "File path not found"
		line := ctx.GetStart().GetLine()

		fmt.Println(r2d2Styles.ErrorMessage(formatFileErrorMessage(msg, fmt.Sprintf("%d", line))))
		v.ErrorCollector.Add(msg, line)

		return nil
	}

	// Extract and clean the file path
	rawPath := ctx.STRING_LITERAL().GetText()
	if rawPath == "\"\"" {
		msg := "Empty file path"
		line := ctx.GetStart().GetLine()

		fmt.Println(r2d2Styles.ErrorMessage(formatFileErrorMessage(msg, fmt.Sprintf("%d", line))))
		v.ErrorCollector.Add(msg, line)

		return nil
	}

	justPath := strings.Trim(rawPath, "\"")

	// Check if the file exists
	if _, err := os.Stat(justPath); err != nil {
		line := ctx.GetStart().GetLine()

		if os.IsNotExist(err) {
			msg := "File not found on path"
			fmt.Println(r2d2Styles.ErrorMessage(formatFileErrorMessage(msg, justPath)))
			v.ErrorCollector.Add(fmt.Sprintf("%s: %s", msg, justPath), line)
		} else {
			msg := "Error checking file"
			fmt.Println(r2d2Styles.ErrorMessage(formatFileErrorMessage(msg, err.Error())))
			v.ErrorCollector.Add(fmt.Sprintf("%s: %s", msg, err.Error()), line)
		}

		return nil
	}

	// Read the content of the file
	content, err := os.ReadFile(justPath)
	if err != nil {
		line := ctx.GetStart().GetLine()
		msg := "Failed to read file"

		fmt.Println(r2d2Styles.ErrorMessage(formatFileErrorMessage(msg, justPath)))
		v.ErrorCollector.Add(fmt.Sprintf("%s: %s", msg, justPath), line)

		return nil
	}

	// Parse the content as if it were part of the same code
	input := antlr.NewInputStream(string(content))
	lexer := parser.NewR2D2Lexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewR2D2Parser(tokens)
	el := parser.NewR2D2ErrorListener(v.ErrorCollector)
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	tree := p.Program()

	// Create a sub-visitor to process the imported content
	sub := NewR2D2Visitor(v.ErrorCollector)
	tree.Accept(sub)

	// Merge the symbol tables
	maps.Copy(v.symbolTable.Modules, sub.symbolTable.Modules)
	maps.Copy(v.symbolTable.Interfaces, sub.symbolTable.Interfaces)
	maps.Copy(v.symbolTable.Globals, sub.symbolTable.Globals)

	// Prepend the imported code
	v.JsCode = sub.JsCode + v.JsCode

	return nil
}

func (v *R2D2Visitor) VisitGlobalDeclaration(ctx *parser.GlobalDeclarationContext) any {

	if ctx.IDENTIFIER() == nil {
		msg := "IDENTIFIER is nil in GlobalDeclaration"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine(msg)))
		v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
		return nil
	}

	if ctx.Expression() == nil {
		msg := "Expression is nil in GlobalDeclaration"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine(msg)))
		v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
		return nil
	}

	if ctx.TypeExpression() == nil {
		msg := "TypeExpression is nil in GlobalDeclaration"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine(msg)))
		v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
		return nil
	}

	globalName := ctx.IDENTIFIER().GetText()

	// Stores in symbolTable
	v.symbolTable.Globals[globalName] = Global{
		Name:  globalName,
		Value: "",
		Type:  ctx.TypeExpression().GetText(),
	}

	//  Generates JS
	v.JsCode += fmt.Sprintf("const %s = ", globalName)
	ctx.Expression().Accept(v)
	v.JsCode += ";"

	return nil
}

func (v *R2D2Visitor) VisitModuleDeclaration(ctx *parser.ModuleDeclarationContext) any {
	// ========================================
	// 1. INITIALIZE SYMBOL TABLES
	// ========================================

	// Guarantee that the global symbol table is initialized
	if v.symbolTable.Globals == nil {
		v.symbolTable.Globals = make(map[string]Global)
	}

	// Ensure that the module symbol table is initialized
	if v.symbolTable.Modules == nil {
		v.symbolTable.Modules = make(map[string]Module)
	}

	// ========================================
	// 2. VALIDATE MODULE IDENTIFIER
	// ========================================

	// Verification to ensure that the module has an identifier
	if ctx.IDENTIFIER(0) == nil || len(ctx.AllIDENTIFIER()) == 0 {
		msg := "Module declaration without identifier"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, ctx.GetStart().GetLine())))
		v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
		return nil
	}

	moduleName := ctx.IDENTIFIER(0).GetText()

	// ========================================
	// 3. INITIALIZE MODULE
	// ========================================

	// Initialize the module
	if _, exists := v.symbolTable.Modules[moduleName]; !exists {
		newModule := Module{
			Name:       moduleName,
			Implements: Interface{},
			Functions:  make(map[string]Function),
			Variables:  make(map[string]Variable),
			Types:      make(map[string]any),
		}

		v.symbolTable.Modules[moduleName] = newModule
		v.currentModule = newModule

		// Make sure the module is added to the symbol table after everything is done
		defer func() {
			v.symbolTable.Modules[v.currentModule.Name] = v.currentModule
		}()
	} else {
		msg := fmt.Sprintf("Module '%s' already exists", moduleName)
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessageNoLine(msg)))
		v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
		return nil
	}

	// ========================================
	// 4. PROCESS MODULE CONTENT
	// ========================================

	// Process the module content to populate the symbol table
	for _, child := range ctx.GetChildren() {

		// --- FUNCTION DECLARATIONS ---
		if funcDecl, ok := child.(*parser.FunctionDeclarationContext); ok {
			if funcDecl.IDENTIFIER() == nil {
				continue
			}

			funcName := funcDecl.IDENTIFIER().GetText()

			// Parse the function arguments
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

			// Store the function in the module
			module := v.symbolTable.Modules[moduleName]
			module.Functions[funcName] = function
			v.symbolTable.Modules[moduleName] = module

			if funcName == "main" && !module.Functions["main"].isExported {
				msg := fmt.Sprintf("Main function needs to be exported on line %d", funcDecl.GetStart().GetLine())
				fmt.Println(r2d2Styles.WarningMessage(msg))
			}

		}

		// --- VARIABLE DECLARATIONS ---
		if varDecl, ok := child.(*parser.VariableDeclarationContext); ok {
			if varDecl.IDENTIFIER() == nil {
				continue
			}

			varName := varDecl.IDENTIFIER().GetText()

			tipoo := ""

			if varDecl.TypeExpression() != nil {
				tipoo = varDecl.TypeExpression().GetText()
			}

			variable := Variable{
				Name:       varName,
				Type:       tipoo,
				isExported: isExported(varDecl),
			}

			// Handle variable assignment
			if varDecl.ASSIGN() != nil && varDecl.Expression() != nil {
				expr := varDecl.Expression()
				variable.Value = expr.GetText()
			} else {
				variable.Value = ""
			}

			// Store the variable in the module
			module := v.symbolTable.Modules[moduleName]
			module.Variables[varName] = variable
			v.symbolTable.Modules[moduleName] = module
		}

		// --- TYPE DECLARATIONS ---
		if typeDecl, ok := child.(*parser.TypeDeclarationContext); ok {
			if typeDecl.IDENTIFIER() == nil {
				continue
			}

			typeName := typeDecl.IDENTIFIER().GetText()

			// Store the type in the module
			module := v.symbolTable.Modules[moduleName]
			module.Types[typeName] = typeDecl
			v.symbolTable.Modules[moduleName] = module
		}
	}

	// ========================================
	// 5. HANDLE INTERFACE IMPLEMENTATION
	// ========================================

	// Check if module implements an interface
	if ctx.IMPLEMENTS() != nil {
		implements := ctx.IDENTIFIER(1).GetText()
		module := v.symbolTable.Modules[moduleName]
		module.Implements = v.symbolTable.Interfaces[implements]

		// Validate that all required functions are implemented
		for _, fn := range module.Implements.Functions {
			if _, exists := module.Functions[fn.Name]; !exists {
				msg := fmt.Sprintf("Function '%s' needs to be implemented in module '%s'", fn.Name, moduleName)
				fmt.Println(r2d2Styles.ErrorMessage(msg))
				v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
				continue
			}

			if fn.isPseudo {
				if !module.Functions[fn.Name].isPseudo {
					msg := fmt.Sprintf("Function '%s' needs to be pseudo in module '%s'", fn.Name, moduleName)
					fmt.Println(r2d2Styles.ErrorMessage(msg))
					v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
					continue
				}
			}

			if fn.isExported {
				if !module.Functions[fn.Name].isExported {
					msg := fmt.Sprintf("Function '%s' needs to be exported in module '%s'", fn.Name, moduleName)
					fmt.Println(r2d2Styles.ErrorMessage(msg))
					v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
					continue
				}
			}
		}

		// Validate that all required variables are implemented
		for _, variable := range module.Implements.Variables {
			if _, exists := module.Variables[variable.Name]; !exists {
				msg := fmt.Sprintf("Variable '%s' needs to be implemented in module '%s'", variable.Name, moduleName)
				fmt.Println(r2d2Styles.ErrorMessage(msg))
				v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
			}
		}

		// // Validate that all required variables are implemented
		// for var := range module.Implements.Variables {
		// 	if _, exists := module.Variables[var.Name]; !exists {
		// 		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Variable '%s' needs to be implemented in module '%s'", varName, moduleName)))
		// 	}
		// }

		// Update the module in the symbol table
		v.symbolTable.Modules[moduleName] = module
	}

	// ========================================
	// 6. GENERATE JAVASCRIPT CODE
	// ========================================

	// Start the module in the JS code
	v.JsCode += fmt.Sprintf("const %s = (function () {", moduleName)

	// Visit the children to process the module content
	result := v.VisitChildren(ctx)

	// Get the exported items for the module return statement
	moduleExports := v.symbolTable.Modules[moduleName].Exports()

	// Finish the module with the return statement
	v.JsCode += fmt.Sprintf("return {%s}; })();", strings.Join(moduleExports, ", "))

	// ========================================
	// 7. EXECUTE MAIN FUNCTION IF EXPORTED
	// ========================================

	// Verify if the module exports a main function and execute it
	module := v.symbolTable.Modules[moduleName]
	if mainFunc, exists := module.Functions["main"]; exists && mainFunc.isExported {
		v.JsCode += fmt.Sprintf("%s.main();", moduleName)
	}

	return result

}

func (v *R2D2Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) any {

	if _, ok := ctx.GetParent().(*parser.InterfaceDeclarationContext); ok {
		return nil
	}

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

				var typee string = ""

				if param.TypeExpression() != nil {
					typee = param.TypeExpression().GetText()
				}

				// Create an Argument object
				argument := Argument{
					Name: paramName,
					Type: typee,
				}

				// Add as both an argument and a variable
				v.currentFunction.Arguments[paramName] = argument

				// Also add as a variable in the function's scope
				v.currentFunction.Variables[paramName] = Variable{
					Name:       paramName,
					Type:       typee,
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
					if functionn, ok := stmtCtx.GetChild(0).(*parser.FunctionCallStatementContext); !ok {
						parentFunctionName := parentFuncDecl.IDENTIFIER().GetText()
						line := stmtCtx.GetStart().GetLine()
						msg := fmt.Sprintf("Statement '%s' not allowed in function '%s'", stmtCtx.GetStart().GetText(), parentFunctionName)
						fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
						v.ErrorCollector.Add(msg, line)
					} else {
						functionName := functionn.FunctionCall().GetText()
						re := regexp.MustCompile(`\([^)]*\)`)
						functionName = re.ReplaceAllString(functionName, "")

						if _, ok := v.currentModule.Functions[functionName]; !ok {

							parentFunctionName := parentFuncDecl.IDENTIFIER().GetText()
							line := stmtCtx.GetStart().GetLine()
							msg := fmt.Sprintf("Function '%s' not allowed in function '%s'", functionName, parentFunctionName)
							fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
							v.ErrorCollector.Add(msg, line)
						}

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

	if _, ok := ctx.GetParent().(*parser.InterfaceDeclarationContext); ok {
		return nil
	}

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
			msg := "Const variable must be assigned a value"
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, ctx.GetStart().GetLine())))
			v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
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
			line := ctx.GetStart().GetLine()
			msg := "Variable cannot be exported in a function"
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
			v.ErrorCollector.Add(msg, line)
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
		line := ctx.GetStart().GetLine()
		msg := fmt.Sprintf("Break statement on %s must be within a loop", r2d2Styles.Bold(fmt.Sprintf("line %d", line)))
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	v.JsCode += "break"
	return nil
}

func (v *R2D2Visitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) any {

	// Verify we're in a loop context
	inLoop := findParent(ctx, (*parser.LoopStatementContext)(nil), (*parser.ForStatementContext)(nil), (*parser.WhileStatementContext)(nil))

	if !inLoop {
		line := ctx.GetStart().GetLine()
		errorMessage := fmt.Sprintf("Continue statement on %s must be within a loop", r2d2Styles.Bold(fmt.Sprintf("line %d", line)))
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(errorMessage, line)))
		v.ErrorCollector.Add(errorMessage, line)
	}

	v.JsCode += "continue"
	return nil
}

func (v *R2D2Visitor) VisitForStatement(ctx *parser.ForStatementContext) any {

	v.JsCode += "for ("

	simpleFor := ctx.SimpleFor()

	bothAssignments := len(simpleFor.AllAssignment()) > 1

	if simpleFor != nil {

		// Initialization
		if simpleFor.VariableDeclaration() != nil {
			simpleFor.VariableDeclaration().Accept(v)

		} else if bothAssignments {
			simpleFor.Assignment(0).Accept(v)
		}

		// Condition
		v.JsCode += "; "

		if simpleFor.Expression() != nil {
			simpleFor.Expression().Accept(v)
		}

		// Update
		v.JsCode += "; "

		if simpleFor.AllAssignment() != nil && len(simpleFor.AllAssignment()) > 0 {
			if bothAssignments {
				if simpleFor.Assignment(1).Accept(v) != nil {
					simpleFor.Assignment(1).Accept(v)
				} else {
					line := simpleFor.GetStart().GetLine()
					msg := "Assignment not found"
					fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
					v.ErrorCollector.Add(msg, line)
				}
			} else {
				if simpleFor.Assignment(0) != nil {
					simpleFor.Assignment(0).Accept(v)
				} else {
					line := simpleFor.GetStart().GetLine()
					msg := "Assignment not found"
					fmt.Print((simpleFor.AllAssignment()))
					fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
					v.ErrorCollector.Add(msg, line)
				}
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
	if ctx.IDENTIFIER(0) == nil {
		return nil
	}

	interfaceName := ctx.IDENTIFIER(0).GetText()

	newInterface := Interface{
		Name:      interfaceName,
		Variables: make(map[string]Variable),
		Functions: make(map[string]Function),
	}

	if ctx.IDENTIFIER(1) != nil {
		// newInterface.Extends = ctx.IDENTIFIER(1).GetText()
		if ctx.IDENTIFIER(1) == nil {
			line := ctx.GetStart().GetLine()
			msg := "Interface extends identifier not found"
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
			v.ErrorCollector.Add(msg, line)
			return nil
		}

		if _, ok := v.symbolTable.Interfaces[ctx.IDENTIFIER(1).GetText()]; !ok {
			msg := fmt.Sprintf("Interface '%s' not found", ctx.IDENTIFIER(1).GetText())
			line := ctx.GetStart().GetLine()
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
			v.ErrorCollector.Add(msg, line)
			return nil
		}

		maps.Copy(newInterface.Variables, v.symbolTable.Interfaces[ctx.IDENTIFIER(1).GetText()].Variables)
		maps.Copy(newInterface.Functions, v.symbolTable.Interfaces[ctx.IDENTIFIER(1).GetText()].Functions)
	}

	for _, varDecl := range ctx.AllVariableDeclaration() {

		if varDecl.IDENTIFIER() == nil {
			continue
		}

		if varDecl.ASSIGN() != nil {
			line := varDecl.GetStart().GetLine()
			message := fmt.Sprintf("Variables within interfaces cannot have value: '%s' must be declared without value", varDecl.IDENTIFIER().GetText())
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(message, line)))
			v.ErrorCollector.Add(message, line)
		}

		varName := varDecl.IDENTIFIER().GetText()
		varType := ""

		if varDecl.TypeExpression() != nil {
			varType = varDecl.TypeExpression().GetText()
		}
		isExported := false
		if varDecl.EXPORT() != nil {
			isExported = true
		}

		newInterface.Variables[varName] = Variable{
			Name:       varName,
			Type:       varType,
			Value:      nil,
			isExported: isExported,
		}
	}

	for _, funcDecl := range ctx.AllFunctionDeclaration() {
		if funcDecl.IDENTIFIER() == nil {
			continue
		}

		if funcDecl.Block() != nil {
			line := funcDecl.GetStart().GetLine()
			message := fmt.Sprintf("Functions within interfaces cannot have a block: '%s' must be declared without block", funcDecl.IDENTIFIER().GetText())
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(message, line)))
			v.ErrorCollector.Add(message, line)
		}

		funcName := funcDecl.IDENTIFIER().GetText()

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

		isExported := false

		if funcDecl.EXPORT() != nil {
			isExported = true
		}

		isPseudo := false

		if funcDecl.PSEUDO() != nil {
			isPseudo = true
		}

		newInterface.Functions[funcName] = Function{
			Name:       funcName,
			Arguments:  arguments,
			Variables:  make(map[string]Variable),
			Functions:  make(map[string]Function),
			isExported: isExported,
			isPseudo:   isPseudo,
		}
	}

	if v.symbolTable.Interfaces == nil {
		fmt.Println("symbolTable.Interfaces nil — inicializando. sus isto n ]e para aparecer se esta msg ")
		v.symbolTable.Interfaces = make(map[string]Interface)
	}

	v.symbolTable.Interfaces[interfaceName] = newInterface

	return nil // Evita visitar os filhos outra vez
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
	return v.VisitChildren(ctx)
}

// VisitAssignment handles assignments
func (v *R2D2Visitor) VisitAssignment(ctx *parser.AssignmentContext) any {

	if ctx.IDENTIFIER() == nil {
		msg := "Assignment identifier missing"
		line := ctx.GetStart().GetLine()
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
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
		line := ctx.GetStart().GetLine()
		errorMessage := fmt.Sprintf("Variable '%s' not declared on line %d", r2d2Styles.Bold(varName), line)
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(errorMessage, line)))

		v.ErrorCollector.Add(errorMessage, line)

		v.JsCode += fmt.Sprintf("/* ERROR: %s */", errorMessage)
		return nil
	}

	// Generate assignment code - start with variable name
	v.JsCode += varName

	// Check if this is an array indexing assignment
	if ctx.LBRACK() != nil && ctx.RBRACK() != nil {

		// Handle array indexing: IDENTIFIER LBRACK expression RBRACK
		v.JsCode += "["

		if expressionList := ctx.AllExpression(); len(expressionList) > 0 {
			// Visit the index expression (first expression)
			expressionList[0].Accept(v)
		}

		v.JsCode += "]"

		// if variabl.Type!= "array" {
		// 	errorMessage := fmt.Sprintf("Cannot use index operator on non-array variable '%s' on line %d",
		// 		r2d2Styles.Bold(varName), ctx.GetStart().GetLine())
		// 	fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(errorMessage, ctx.GetStart().GetLine())))
		// 	v.JsCode += fmt.Sprintf("/* ERROR: %s */", errorMessage)
		// }
	}

	// Handle assignment operator or increment/decrement
	if ctx.AssignmentOperator() != nil {

		if ctx.AssignmentOperator().GetText() != "=" {

			isParam, _ := isParameter(v, ctx.IDENTIFIER().GetText())

			if (variabl.Value == nil || variabl.Value == "") && !isParam {
				line := ctx.GetStart().GetLine()
				msg := fmt.Sprintf("Variable '%s' not initialized on %s", ctx.IDENTIFIER().GetText(), r2d2Styles.Bold(fmt.Sprintf("line %d", line)))
				fmt.Println(r2d2Styles.ErrorMessage(msg))
				v.ErrorCollector.Add(msg, line)
			}
		}

		v.JsCode += " " + ctx.AssignmentOperator().GetText() + " "

		// Get the right expression
		// For array indexing, this would be the second expression
		expressionList := ctx.AllExpression()

		if ctx.LBRACK() != nil && len(expressionList) > 1 {

			// Visit the value expression (second expression)
			expressionList[1].Accept(v)

		} else if ctx.LBRACK() == nil && len(expressionList) > 0 {

			// Regular assignment with just one expression
			expressionList[0].Accept(v)
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
		line := ctx.GetStart().GetLine()
		msg := "Missing function identifier"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
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
		line := ctx.GetStart().GetLine()
		msg := fmt.Sprintf("Function '%s' not found", funcName)
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
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
			var msg string
			skips := false
			switch len(exprs) > len(fn.Arguments) {
			case true:

				for _, arg := range fn.Arguments {
					if arg.isInfinite() {
						skips = true
						break
					}
				}
				if !skips {
					msg = fmt.Sprintf(
						"Too many arguments passed when calling function %s",
						r2d2Styles.Bold(funcName),
					)
				}

			case false:

				msg = fmt.Sprintf(
					"Too little arguments passed when calling function %s",
					r2d2Styles.Bold(funcName),
				)
			}

			if !skips {
				fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, ctx.GetStart().GetLine())))
				v.ErrorCollector.Add(msg, ctx.GetStart().GetLine())
			}
		}

		for i, expr := range exprs {
			expr.Accept(v) // Visits normally the argument

			if i < len(exprs)-1 {
				v.JsCode += ", "
			}
		}
	}

	v.JsCode += ")"

	return nil
}

func (v *R2D2Visitor) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) any {

	text := ctx.Literal().GetText()

	litCtx := ctx.Literal()

	for i := range litCtx.GetChildren() {
		child := litCtx.GetChild(i)
		if objLit, ok := child.(*parser.ObjectLiteralContext); ok {
			// Encontraste objectLiteral, visita-o!
			return v.VisitObjectLiteral(objLit)
		}
		if arrLit, ok := child.(*parser.ArrayLiteralContext); ok {
			// Se quiseres fazer algo para arrayLiteral
			return v.VisitArrayLiteral(arrLit)
		}
	}

	if strings.HasPrefix(text, `"""`) && strings.HasSuffix(text, `"""`) {
		// Removes the triple quotes and uses the equivalent sintax in JS `
		inner := text[3 : len(text)-3]
		v.JsCode += "`" + inner + "`"
	} else {
		// Normal string
		v.JsCode += text
	}

	return nil
}

func (v *R2D2Visitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) any {

	line := ctx.GetStart().GetLine()

	if ctx.IDENTIFIER() != nil {

		if ok, variable := isAccessibleVariable(v, ctx.IDENTIFIER().GetText()); ok {

			isParam, _ := isParameter(v, ctx.IDENTIFIER().GetText())

			if (variable.Value == nil || variable.Value == "") && !isParam {
				line := ctx.GetStart().GetLine()
				msg := fmt.Sprintf("Variable '%s' not initialized on %s", ctx.IDENTIFIER().GetText(), r2d2Styles.Bold(fmt.Sprintf("line %d", line)))
				fmt.Println(r2d2Styles.ErrorMessage(msg))
				v.ErrorCollector.Add(msg, line)
			}

			v.JsCode += ctx.IDENTIFIER().GetText()

		} else {
			line := ctx.GetStart().GetLine()
			msg := fmt.Sprintf("Variable '%s' not found", ctx.IDENTIFIER().GetText())
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
			v.ErrorCollector.Add(msg, line)
		}
	} else {
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage("Identifier not found", line)))
		v.ErrorCollector.Add("Identifier not found", line)
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
			line := ctx.GetStart().GetLine()
			msg := "Unknown unary operator"
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
			v.ErrorCollector.Add(msg, line)
		}

	} else {
		line := ctx.GetStart().GetLine()
		msg := "Expected a unary operator"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	if ctx.Expression() != nil {
		ctx.Expression().Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		msg := "Expression not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	return nil
}

func (v *R2D2Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) any {

	if ctx.Expression(0) != nil {
		ctx.Expression(0).Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		msg := "Left expression not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
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
			line := ctx.GetStart().GetLine()
			msg := "Unknown multiplicative operator"
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
			v.ErrorCollector.Add(msg, line)
		}

	} else {
		line := ctx.GetStart().GetLine()
		msg := "Expected a multiplicative operator"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	if ctx.Expression(1) != nil {
		ctx.Expression(1).Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		msg := "Right expression not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	return nil
}

func (v *R2D2Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) any {

	if ctx.Expression(0) != nil {
		ctx.Expression(0).Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		msg := "Left expression not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
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
			line := ctx.GetStart().GetLine()
			msg := "Unknown additive operator"
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
			v.ErrorCollector.Add(msg, line)
		}

	} else {
		line := ctx.GetStart().GetLine()
		msg := "Expected an additive operator"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	if ctx.Expression(1) != nil {
		ctx.Expression(1).Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		msg := "Right expression not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	return nil
}

func (v *R2D2Visitor) VisitComparisonExpression(ctx *parser.ComparisonExpressionContext) any {

	if ctx.Expression(0) != nil {
		ctx.Expression(0).Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		msg := "Left expression not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
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
			line := ctx.GetStart().GetLine()
			msg := "Unknown comparison operator"
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
			v.ErrorCollector.Add(msg, line)
		}

	} else {
		line := ctx.GetStart().GetLine()
		msg := "Expected a comparison operator"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	if ctx.Expression(1) != nil {
		ctx.Expression(1).Accept(v)
	} else {

		line := ctx.GetStart().GetLine()
		msg := "Right expression not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	return nil
}

func (v *R2D2Visitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) any {

	if ctx.Expression(0) != nil {
		ctx.Expression(0).Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		msg := "Left expression not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	if opNode, ok := ctx.GetChild(1).(antlr.TerminalNode); ok {

		token := opNode.GetSymbol().GetTokenType()

		switch token {
		case parser.R2D2ParserAND:
			v.JsCode += "&&"
		case parser.R2D2ParserOR:
			v.JsCode += "||"
		default:
			line := ctx.GetStart().GetLine()
			msg := "Unknown logical operator"
			fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
			v.ErrorCollector.Add(msg, line)
		}

	} else {
		line := ctx.GetStart().GetLine()
		msg := "Expected a logical operator"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	if ctx.Expression(1) != nil {
		ctx.Expression(1).Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		msg := "Right expression not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	return nil
}

func (v *R2D2Visitor) VisitJsStatement(ctx *parser.JsStatementContext) any {

	if ctx.STRING_LITERAL() != nil {
		raw := ctx.STRING_LITERAL().GetText()

		if strings.HasPrefix(raw, `"""`) && strings.HasSuffix(raw, `"""`) {

			clean := strings.TrimSuffix(strings.TrimPrefix(raw, `"""`), `"""`)
			v.JsCode += clean

		} else if strings.HasPrefix(raw, `"`) && strings.HasSuffix(raw, `"`) {

			clean := strings.TrimSuffix(strings.TrimPrefix(raw, `"`), `"`)
			v.JsCode += clean

		} else {
			v.JsCode += raw
		}
	} else {
		line := ctx.GetStart().GetLine()
		msg := "JS code not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	return nil
}

func (v *R2D2Visitor) VisitArrayAccessExpression(ctx *parser.ArrayAccessExpressionContext) any {

	if ctx.Expression(0) != nil {
		ctx.Expression(0).Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		msg := "Array not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	if ctx.LBRACK() != nil {
		v.JsCode += "["
	}

	if ctx.Expression(1) != nil {
		ctx.Expression(1).Accept(v)
	} else {
		line := ctx.GetStart().GetLine()
		msg := "Index not found"
		fmt.Println(r2d2Styles.ErrorMessage(formatErrorMessage(msg, line)))
		v.ErrorCollector.Add(msg, line)
	}

	if ctx.RBRACK() != nil {
		v.JsCode += "]"
	}

	return nil
}

func (v *R2D2Visitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) any {
	v.JsCode += "["

	// Use AllExpression() to get the slice of expressions
	for i, expr := range ctx.AllExpression() {
		if i > 0 {
			v.JsCode += ", "
		}
		expr.Accept(v)
	}

	v.JsCode += "]"
	return nil
}

func (v *R2D2Visitor) VisitObjectLiteral(ctx *parser.ObjectLiteralContext) any {

	v.JsCode += "{"
	for i, expr := range ctx.AllExpression() {
		if (i % 2) == 0 {
			expr.Accept(v)
			v.JsCode += ": "
			continue
		}

		expr.Accept(v) // Visits normally the argument
		v.JsCode += ", "
	}

	v.JsCode += "}"

	return nil
}
