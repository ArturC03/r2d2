package visitor

import "github.com/ArturC03/r2d2/parser"

// Represents a R2D2 variable
type Variable struct {
	Name       string
	Value      any
	Type       string
	isExported bool
}

// Represents a R2D2 global variable - doesnt really exist
type Global struct {
	Name  string
	Value any
	Type  string
}

// Represents a R2D2 function argument/parameter whatever you want to call it
type Argument struct {
	Name string
	Type string
}

// Represents a R2D2 function
type Function struct {
	Name       string
	Arguments  map[string]Argument
	Variables  map[string]Variable
	Functions  map[string]Function
	isExported bool
	isPseudo   bool
}

// Represents a R2D2 module
type Module struct {
	Name       string
	Implements Interface
	Functions  map[string]Function
	Variables  map[string]Variable
	Types      map[string]any
}

// Represents a R2D2 interface - TODO: didnt really implement it, doing it now
type Interface struct {
	Name      string
	Variables map[string]Variable
	Functions map[string]Function
}

// Represents the symbol table for the R2D2 visitor
type SymbolTable struct {
	Modules    map[string]Module
	Interfaces map[string]Interface
	Globals    map[string]Global
}

// Represents the R2D2 visitor structure
type R2D2Visitor struct {
	parser.BaseR2D2Visitor
	symbolTable      SymbolTable
	JsCode           string
	currentModule    Module
	currentFunction  Function
	currentInterface Interface
}
