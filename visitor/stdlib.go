package visitor

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"os/exec"
	"strings"

	r2d2Styles "github.com/ArturC03/r2d2Styles"
)

//go:embed ts/list_std_functions.ts
var denoScript string

// Returns a map of all the standard library functions
func GetStdFunctions() map[string]Module {
	// Criar ficheiro temporário com o conteúdo embutido
	tmpFile, err := os.CreateTemp("", "list_std_functions_*.ts")
	if err != nil {
		panic("Cannot create temp file: " + err.Error())
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(denoScript)
	if err != nil {
		panic("Cannot write to temp file: " + err.Error())
	}
	tmpFile.Close()

	// Executar o script com Deno
	cmd := exec.Command("deno", "run", "--allow-all", tmpFile.Name())

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage("Failed to run Deno: " + err.Error()))
		return nil
	}

	// Parse do JSON
	var raw map[string][]string
	err = json.Unmarshal(out, &raw)
	if err != nil {
		panic("Invalid JSON from Deno: " + err.Error())
	}

	modules := make(map[string]Module)

	for fullName, args := range raw {
		parts := strings.Split(fullName, ".")
		modName := "std"
		funcName := fullName

		if len(parts) > 1 {
			modName = parts[0]
			funcName = parts[1]
		}

		if _, ok := modules[modName]; !ok {
			modules[modName] = Module{
				Name:      modName,
				Functions: make(map[string]Function),
				Variables: make(map[string]Variable),
				Types:     make(map[string]any),
			}
		}

		argMap := make(map[string]Argument)
		for _, arg := range args {
			argMap[arg] = Argument{Name: arg, Type: "any"}
		}

		fn := Function{
			Name:       funcName,
			Arguments:  argMap,
			isExported: true,
			isPseudo:   true,
		}

		mod := modules[modName]
		mod.Functions[funcName] = fn
		modules[modName] = mod
	}

	return modules
}

// Loads the standard library functions from Deno into multiple std modules
func (v *R2D2Visitor) LoadStdModules() {
	if v.symbolTable.Modules == nil {
		v.symbolTable.Modules = make(map[string]Module)
	}

	stdModules := GetStdFunctions()

	for name, mod := range stdModules {
		v.symbolTable.Modules[name] = mod
		maps.Copy(v.symbolTable.Modules[name].Functions, mod.Functions)
	}
}
