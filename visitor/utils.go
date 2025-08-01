package visitor

import (
	"github.com/ArturC03/r2d2/errors"
	"github.com/ArturC03/r2d2/parser"
	"github.com/antlr4-go/antlr/v4"

	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ArturC03/r2d2Styles"
)

// Returns a list of the exported function and variable names for a given module
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

// Creates a new R2D2 visitor instance
func NewR2D2Visitor(errorCollector *errors.ErrorCollector) *R2D2Visitor {
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
		currentFunction: Function{
			Variables: make(map[string]Variable),
			Functions: make(map[string]Function),
		},
		currentInterface: Interface{},
		JsCode:           "",
		ErrorCollector:   errorCollector,
	}
}

// Checks if a node is exported
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

// Checks if a function is a pseudo function - aka imposter
func isPseudo(node *parser.FunctionDeclarationContext) bool {
	return node.PSEUDO() != nil
}

// Finds the first child of a node that matches any of the target types
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

// Finds the parent of a node that matches any of the target types
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
		}
	}

	// Recursively check parent's parent
	return findParent(parent.(antlr.RuleContext), types...)
}

// Helper function to check if the passed argument matches the expected type
func isValidJSType(passedArg string, expectedType string) bool {
	// Accept anything for "any" type
	if expectedType == "any" {
		return true
	}

	// Check for number type
	if expectedType == "number" {

		// Exclude strings and booleans
		if strings.Contains(passedArg, "\"") || strings.Contains(passedArg, "'") {
			return false // Is a string
		}

		if passedArg == "true" || passedArg == "false" {
			return false // Is a boolean
		}

		// Try to parse as number
		if _, err := fmt.Sscanf(passedArg, "%f", new(float64)); err == nil {
			return true
		}

		return false
	}

	// Check for string type
	if expectedType == "string" {
		return strings.HasPrefix(passedArg, "\"") || strings.HasPrefix(passedArg, "'")
	}

	// Check for boolean type
	if expectedType == "boolean" {
		return passedArg == "true" || passedArg == "false"
	}

	// For objects, arrays, and other complex types
	// Accept by default as JavaScript is flexible with types
	return true
}

// Function to load all global JavaScript functions from Deno
func loadGlobalFunctions(v *R2D2Visitor) error {
	// Check if global functions are already loaded
	if module, exists := v.symbolTable.Modules["global"]; exists && len(module.Functions) > 0 {
		return nil
	}

	// Initialize global module if it doesn't exist
	if _, exists := v.symbolTable.Modules["global"]; !exists {
		v.symbolTable.Modules["global"] = Module{
			Name:      "global",
			Functions: make(map[string]Function),
			Variables: make(map[string]Variable),
			Types:     make(map[string]any),
		}
	}

	// Also make sure standard modules like console and Math exist
	standardModules := []string{"console", "Math", "Array", "String", "Object", "Date"}
	for _, modName := range standardModules {
		if _, exists := v.symbolTable.Modules[modName]; !exists {
			v.symbolTable.Modules[modName] = Module{
				Name:      modName,
				Functions: make(map[string]Function),
				Variables: make(map[string]Variable),
				Types:     make(map[string]any),
			}
		}
	}

	// Deno command to discover global functions and their signatures
	cmd := exec.Command("deno", "eval", `
	const availableFunctions = {};

				// Helper to extract parameter names and types from function
				function getParameterInfo(func) {
								try {
												const funcStr = func.toString();
												const params = [];

												// Get parameters from function string
												const match = funcStr.match(/\\((.*?)\\)/);
												if (match && match[1].trim()) {
																const paramNames = match[1].split(',').map(p => p.trim());

																paramNames.forEach((param, index) => {
																				// Remove any default values
																				const cleanParam = param.split('=')[0].trim();
																				// Remove any type annotations
																				const finalParam = cleanParam.split(':')[0].trim();

																				params.push({
																								name: finalParam || 'arg' + (index + 1),
																								type: inferParameterType(func, index)
																				});
																});
												}
												return params;
								} catch (e) {
												return [];
								}
				}

				// Helper to infer parameter type by testing the function
				function inferParameterType(func, paramIndex) {
								try {
												// Create test values of different types
												const testValues = {
																'number': 42,
																'string': 'test',
																'boolean': true,
																'object': {},
																'array': [],
																'function': () => {}
												};

												// Try to call the function with different types
												for (const [type, value] of Object.entries(testValues)) {
																const args = Array(paramIndex).fill(undefined);
																args.push(value);

																try {
																				func.apply(null, args);
																				return type;
																} catch (e) {
																				// If error contains type information, use it
																				if (e instanceof TypeError) {
																								const errorMsg = e.toString().toLowerCase();
																								if (errorMsg.includes('number')) return 'number';
																								if (errorMsg.includes('string')) return 'string';
																								if (errorMsg.includes('boolean')) return 'boolean';
																								if (errorMsg.includes('function')) return 'function';
																								if (errorMsg.includes('object')) return 'object';
																				}
																}
												}
								} catch (e) {}

								return 'any'; // Default to 'any' if type cannot be determined
				}

				// Process global objects and their methods
				function processObject(obj, name, isPrototype = false) {
								if (obj === null || obj === undefined) return;

								try {
												const props = Object.getOwnPropertyNames(obj);
												props.forEach(prop => {
																try {
																				const fullName = isPrototype ? name + '.prototype.' + prop : name + '.' + prop;
																				const descriptor = Object.getOwnPropertyDescriptor(obj, prop);

																				if (descriptor && typeof descriptor.value === 'function') {
																								availableFunctions[fullName] = {
																												type: 'function',
																												parameters: getParameterInfo(descriptor.value),
																												returnType: 'any'
																								};
																				}
																} catch (e) {}
												});
								} catch (e) {}
				}

				// Get all global objects
				const globals = Object.getOwnPropertyNames(globalThis);
				globals.forEach(name => {
								try {
												const obj = globalThis[name];

												if (typeof obj === 'function') {
																// Store function itself
																availableFunctions[name] = {
																				type: 'function',
																				parameters: getParameterInfo(obj),
																				returnType: 'any'
																};

																// Process prototype methods
																processObject(obj.prototype, name, true);
												} else if (typeof obj === 'object' && obj !== null) {
																// Process object methods
																processObject(obj, name);
												}
								} catch (e) {}
				});

				console.log(JSON.stringify(availableFunctions, null, 2));
				`)

	// Execute command and get output
	output, err := cmd.CombinedOutput()
	if err != nil {
		errMsg := fmt.Sprintf("Error executing Deno command: %v", err)
		fmt.Println(r2d2Styles.ErrorMessage(errMsg))
		v.ErrorCollector.Add(errMsg, 0)
		return fmt.Errorf("Error executing Deno command: %w", err)
	}

	// Parse JSON output
	var result map[string]any
	if err := json.Unmarshal(output, &result); err != nil {
		line := 1 // Se não houver linha específica, usa 1 ou outra linha padrão
		lineMsg := r2d2Styles.Bold(fmt.Sprintf("Error parsing JSON at line %d", line))
		errMsg := fmt.Sprintf("Error parsing JSON: %v. %s", err, lineMsg)
		fmt.Println(r2d2Styles.ErrorMessage(errMsg))
		if v.ErrorCollector != nil {
			v.ErrorCollector.Add(errMsg, line)
		}
		return fmt.Errorf("Error parsing JSON: %w", err)
	}

	globalModule := v.symbolTable.Modules["global"]
	count := 0

	// Process discovered functions
	for funcName, info := range result {
		infoMap, ok := info.(map[string]any)
		if !ok {
			continue
		}

		// Skip if not a function
		funcType, ok := infoMap["type"].(string)
		if !ok || funcType != "function" {
			continue
		}

		// Get parameters
		var args = make(map[string]Argument)
		params, ok := infoMap["parameters"].([]interface{})

		if ok {
			for i, param := range params {

				paramMap, ok := param.(map[string]interface{})
				if !ok {
					continue
				}

				name, _ := paramMap["name"].(string)
				paramType, _ := paramMap["type"].(string)

				// Ensure name isn't empty
				if name == "" {
					name = fmt.Sprintf("arg%d", i+1)
				}

				args[name] = Argument{
					Name: name,
					Type: paramType,
				}
			}
		}

		// Determine if this is a module method
		parts := strings.Split(funcName, ".")
		if len(parts) >= 2 {
			moduleName := parts[0]
			methodName := parts[1]

			// Skip prototype methods for simplicity
			if methodName == "prototype" {
				continue
			}

			// Use proper module if it's a known standard module
			if module, exists := v.symbolTable.Modules[moduleName]; exists {
				module.Functions[methodName] = Function{
					Name:       methodName,
					Arguments:  args,
					Variables:  make(map[string]Variable),
					Functions:  make(map[string]Function),
					isExported: true,
				}
				v.symbolTable.Modules[moduleName] = module
			} else {
				// Store in global module with full name
				globalModule.Functions[funcName] = Function{
					Name:       funcName,
					Arguments:  args,
					Variables:  make(map[string]Variable),
					Functions:  make(map[string]Function),
					isExported: true,
				}
			}
		} else {
			// Store global function
			globalModule.Functions[funcName] = Function{
				Name:       funcName,
				Arguments:  args,
				Variables:  make(map[string]Variable),
				Functions:  make(map[string]Function),
				isExported: true,
			}
		}
		count++
	}

	// Add standard library functions explicitly if they weren't discovered
	standardFunctions := map[string]map[string]Argument{
		"console.log": {
			"message": {Name: "message", Type: "any"},
		},
		"console.error": {
			"message": {Name: "message", Type: "any"},
		},
		"console.warn": {
			"message": {Name: "message", Type: "any"},
		},
		"Math.random": {},
		"Math.floor": {
			"value": {Name: "value", Type: "number"},
		},
	}

	// Add standard functions to appropriate modules
	for fullName, args := range standardFunctions {
		parts := strings.Split(fullName, ".")
		if len(parts) == 2 {
			moduleName, funcName := parts[0], parts[1]

			if module, exists := v.symbolTable.Modules[moduleName]; exists {
				if _, funcExists := module.Functions[funcName]; !funcExists {
					module.Functions[funcName] = Function{
						Name:       funcName,
						Arguments:  args,
						Variables:  make(map[string]Variable),
						Functions:  make(map[string]Function),
						isExported: true,
					}
					v.symbolTable.Modules[moduleName] = module
					count++
				}
			}
		}
	}

	// Update global module
	v.symbolTable.Modules["global"] = globalModule

	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("Loaded %d global functions", count)))
	return nil
}

// Function to find similar functions by name
func findSimilarFunctions(functions map[string]Function, name string) []string {
	var suggestions []string

	// If the name contains a dot (object.method)
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		objName, _ := parts[0], parts[1] // objName, methodName

		// Search for methods of the same object
		for funcName := range functions {
			if strings.HasPrefix(funcName, objName+".") {
				suggestions = append(suggestions, funcName)
				if len(suggestions) >= 3 {
					break
				}
			}
		}
	} else {
		// Search for top-level functions or objects with similar names
		for funcName := range functions {
			// If it's a top-level function (no dot)
			if !strings.Contains(funcName, ".") {
				if len(name) > 0 && len(funcName) > 0 &&
					(strings.HasPrefix(funcName, name[:1]) ||
						levenshteinDistance(funcName, name) <= 3) {
					suggestions = append(suggestions, funcName)
				}
			} else if !strings.Contains(name, ".") {
				// If we're looking for a name that could be an object
				objName := strings.Split(funcName, ".")[0]
				if objName == name {
					suggestions = append(suggestions, funcName)
				}
			}

			if len(suggestions) >= 3 {
				break
			}
		}
	}

	return suggestions
}

// Function to check if a function is accessible and valid
func (v *R2D2Visitor) isAccessibleFunction(funcName string) (bool, Function, string) {
	parts := strings.Split(funcName, ".")

	if len(parts) == 2 {
		// Module.function format
		moduleName, methodName := parts[0], parts[1]

		// Check if module exists
		module, moduleExists := v.symbolTable.Modules[moduleName]
		if !moduleExists {
			// Try to find similar modules/functions
			globalModule, globalExists := v.symbolTable.Modules["global"]
			if !globalExists {
				return false, Function{}, "/* ERROR: Global module not initialized */"
			}

			suggestions := findSimilarFunctions(globalModule.Functions, funcName)
			errorMsg := fmt.Sprintf("/* ERROR: Module '%s' not found */", moduleName)
			if len(suggestions) > 0 {
				errorMsg += fmt.Sprintf(" /* Você quis dizer: %s? */", strings.Join(suggestions, ", "))
			}
			return false, Function{}, errorMsg
		}

		// Check if function exists in module
		function, functionExists := module.Functions[methodName]
		if !functionExists {
			suggestions := findSimilarFunctions(module.Functions, methodName)
			errorMsg := fmt.Sprintf("/* ERROR: Function '%s' not found in module '%s' */", methodName, moduleName)
			if len(suggestions) > 0 {
				errorMsg += fmt.Sprintf(" /* Você quis dizer: %s? */", strings.Join(suggestions, ", "))
			}
			return false, Function{}, errorMsg
		}

		// Check if function is exported (if calling from another module)
		if moduleName != v.currentModule.Name && !function.isExported {
			return false, Function{}, fmt.Sprintf("/* ERROR: Function '%s' is not exported from module '%s' */", methodName, moduleName)
		}

		return true, function, ""
	} else {
		// Check current module functions
		if v.currentModule.Name != "" {
			if function, exists := v.currentModule.Functions[funcName]; exists {
				return true, function, ""
			}
		}

		// Check global module
		if globalModule, exists := v.symbolTable.Modules["global"]; exists {
			if function, exists := globalModule.Functions[funcName]; exists {
				return true, function, ""
			}
		}

		// Check global variables that could be functions
		if _, exists := v.symbolTable.Globals[funcName]; exists {
			// Convert Global to Function for interface compatibility
			return true, Function{
				Name:       funcName,
				Arguments:  make(map[string]Argument),
				isExported: true,
			}, ""
		}

		// Function not found - suggest similar functions
		globalModule, exists := v.symbolTable.Modules["global"]
		if !exists {
			return false, Function{}, fmt.Sprintf("/* ERROR: Function '%s' not found */", funcName)
		}

		suggestions := findSimilarFunctions(globalModule.Functions, funcName)
		errorMsg := fmt.Sprintf("/* ERROR: Function '%s' not found */", funcName)
		if len(suggestions) > 0 {
			errorMsg += fmt.Sprintf(" /* Você quis dizer: %s? */", strings.Join(suggestions, ", "))
		}
		return false, Function{}, errorMsg
	}
}

// Levenshtein distance implementation for finding similar names
func levenshteinDistance(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	// Initialize distance matrix
	matrix := make([][]int, len(a)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
		matrix[i][0] = i
	}
	for j := range matrix[0] {
		matrix[0][j] = j
	}

	// Calculate distances
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deletion
				matrix[i][j-1]+1,      // insertion
				matrix[i-1][j-1]+cost, // substitution
			)
		}
	}

	return matrix[len(a)][len(b)]
}

// Helper function to find the minimum value among three integers
func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

func isAccessibleVariable(v *R2D2Visitor, varName string) (bool, Variable) {
	// Check if variable is a global-level variable
	if _, exists := v.symbolTable.Globals[varName]; exists {
		return true, Variable{}
	}

	// Check if variable is a module-level variable
	if _, exists := v.currentModule.Variables[varName]; exists {
		return true, v.currentModule.Variables[varName]
	}

	// Check if variable is a function-level variable
	if _, exists := v.currentFunction.Variables[varName]; exists {
		return true, v.currentFunction.Variables[varName]
	}

	return false, Variable{}

}

func isAccessibleFunction(v *R2D2Visitor, funcName string) (bool, Function) {
	// Split function name by module prefix if it exists
	parts := strings.Split(funcName, ".")

	if len(parts) == 2 {
		// Module.function format
		moduleName, methodName := parts[0], parts[1]

		// Check if module exists
		if module, exists := v.symbolTable.Modules[moduleName]; exists {
			// Check if function exists in module
			if fn, exists := module.Functions[methodName]; exists && fn.isExported {
				return true, fn
			}
		}
		return false, Function{}
	}

	// Check if function is a module-level function
	if _, exists := v.currentModule.Functions[funcName]; exists {
		return true, v.currentModule.Functions[funcName]
	}

	// Check if function is a function-level function
	if _, exists := v.currentFunction.Functions[funcName]; exists {

		return true, v.currentFunction.Functions[funcName]
	}

	// Check all modules for exported functions
	// for _, module := range v.symbolTable.Modules {
	// 	if fn, exists := module.Functions[funcName]; exists && fn.isExported {
	// 		return true, fn
	// 	}
	// }

	return false, Function{}
}

// Checks if a variable is a parameter of the current function
func isParameter(v *R2D2Visitor, varName string) (bool, Argument) {
	if fn, exists := v.currentFunction.Arguments[varName]; exists {
		return true, fn
	}

	return false, Argument{}
}

func (v *R2D2Visitor) RegisterInterfaceNames(ctx *parser.ProgramContext) {
	for _, ifaceCtx := range ctx.AllInterfaceDeclaration() {
		if ifaceCtx.IDENTIFIER(0) == nil {
			continue
		}
		interfaceName := ifaceCtx.IDENTIFIER(0).GetText()

		if v.symbolTable.Interfaces == nil {
			v.symbolTable.Interfaces = make(map[string]Interface)
		}

		// Só regista o nome e estruturas vazias
		if _, exists := v.symbolTable.Interfaces[interfaceName]; !exists {
			v.symbolTable.Interfaces[interfaceName] = Interface{
				Name:      interfaceName,
				Variables: make(map[string]Variable),
				Functions: make(map[string]Function),
			}
		}
	}
}

func (arg Argument) isInfinite() bool {
	if strings.Contains(arg.Name, "...") {
		return true
	}

	return false
}
