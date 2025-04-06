package visitor

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	// "github.com/ArturC03/r2d2/parser"
	r2d2Styles "github.com/ArturC03/r2d2Styles"
)

// Função para carregar as funções globais do JavaScript
// Esta função recebe diretamente o visitante como parâmetro em vez de usar uma variável global
func loadGlobalFunctions(v *R2D2Visitor) error {
	// Verificar se o módulo global já existe e já tem funções carregadas
	if module, exists := v.symbolTable.Modules["global"]; exists && len(module.Functions) > 0 {
		return nil // Funções já carregadas
	}

	// Inicializar o módulo global se não existir
	if _, exists := v.symbolTable.Modules["global"]; !exists {
		v.symbolTable.Modules["global"] = Module{
			Name:      "global",
			Functions: make(map[string]Function),
			Variables: make(map[string]Variable),
			Types:     make(map[string]any),
		}
	}

	// Comando Deno para obter as funções globais com suas assinaturas
	cmd := exec.Command("deno", "eval", `
const availableFunctions = {};

// Função para extrair os nomes dos parâmetros
function getParameterNames(func) {
  try {
    const funcStr = func.toString();
    const match = funcStr.match(/\\(([^)]+)\\)/);
    if (match && match[1].trim()) {
      return match[1].split(',').map(param => param.trim());
    }
  } catch (e) {}
  return [];
}

// Analisar objetos globais
// Primeiro coletamos objetos de alto nível
const highLevelObjects = new Set(Object.getOwnPropertyNames(globalThis));

// Processamos cada objeto de alto nível
highLevelObjects.forEach(name => {
  try {
    const obj = globalThis[name];
    
    // Para objetos que são funções
    if (typeof obj === 'function') {
      try {
        availableFunctions[name] = {
          type: 'function',
          parameterNames: getParameterNames(obj),
          returnType: 'any'
        };
      } catch (e) {
        availableFunctions[name] = {
          type: 'function',
          parameterNames: [],
          returnType: 'any'
        };
      }
    } 
    // Para objetos regulares (como console, Math, etc.)
    else if (typeof obj === 'object' && obj !== null) {
      availableFunctions[name] = {
        type: 'object',
        methods: {}
      };
      
      // Tentar obter os métodos do objeto
      try {
        const props = Object.getOwnPropertyNames(obj);
        props.forEach(prop => {
          try {
            const method = obj[prop];
            if (typeof method === 'function') {
              availableFunctions[name].methods[prop] = {
                type: 'function',
                parameterNames: getParameterNames(method),
                returnType: 'any'
              };
            }
          } catch (e) {
            // Ignorar propriedades que não podem ser acessadas
          }
        });
      } catch (e) {
        // Ignorar objetos que não permitem enumeração de propriedades
      }
    }
  } catch (e) {
    // Ignorar objetos inacessíveis
  }
});

console.log(JSON.stringify(availableFunctions, null, 2));
`)

	// Capturar a saída do comando
	output, err := cmd.CombinedOutput()
	if err != nil {
		errMsg := fmt.Sprintf("/* ERROR: Erro ao executar o comando Deno: %v */", err)
		fmt.Println(r2d2Styles.ErrorMessage(errMsg))
		return fmt.Errorf(errMsg)
	}

	// Fazer o parse da saída JSON
	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		errMsg := fmt.Sprintf("/* ERROR: Erro ao parsear o JSON: %v */", err)
		fmt.Println(r2d2Styles.ErrorMessage(errMsg))
		return fmt.Errorf(errMsg)
	}

	globalModule := v.symbolTable.Modules["global"]
	count := 0

	// Processar funções e objetos globais
	for objName, value := range result {
		valueMap, ok := value.(map[string]interface{})
		if !ok {
			continue
		}

		objType, _ := valueMap["type"].(string)

		if objType == "function" {
			// Processar função global
			paramNamesInterface, _ := valueMap["parameterNames"].([]interface{})
			var args []Argument

			for i, paramNameInterface := range paramNamesInterface {
				paramStr, ok := paramNameInterface.(string)
				if !ok {
					paramStr = fmt.Sprintf("arg%d", i+1)
				}

				args = append(args, Argument{
					Name: paramStr,
					Type: "any",
				})
			}

			globalModule.Functions[objName] = Function{
				Name:       objName,
				Arguments:  args,
				Variables:  make(map[string]Variable),
				Functions:  make(map[string]Function),
				isExported: true,
				isPseudo:   false,
			}

			count++
		} else if objType == "object" {
			// Processar objeto com métodos
			methods, ok := valueMap["methods"].(map[string]interface{})
			if !ok {
				continue
			}

			for methodName, methodInfo := range methods {
				methodMap, ok := methodInfo.(map[string]interface{})
				if !ok {
					continue
				}

				paramNamesInterface, _ := methodMap["parameterNames"].([]interface{})
				var args []Argument

				for i, paramNameInterface := range paramNamesInterface {
					paramStr, ok := paramNameInterface.(string)
					if !ok {
						paramStr = fmt.Sprintf("arg%d", i+1)
					}

					args = append(args, Argument{
						Name: paramStr,
						Type: "any",
					})
				}

				fullName := objName + "." + methodName
				globalModule.Functions[fullName] = Function{
					Name:       fullName,
					Arguments:  args,
					Variables:  make(map[string]Variable),
					Functions:  make(map[string]Function),
					isExported: true,
					isPseudo:   false,
				}

				count++
			}
		}
	}

	// Atualizar o módulo na tabela de símbolos
	v.symbolTable.Modules["global"] = globalModule

	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("/* SUCCESS: Carregadas %d funções globais */", count)))
	return nil
}

// Função para encontrar funções com nomes similares
func findSimilarFunctions(functions map[string]Function, name string) []string {
	var suggestions []string

	// Se o nome contém um ponto (objeto.método)
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		objName, _ := parts[0], parts[1] // objName, methodName

		// Buscar métodos do mesmo objeto
		for funcName := range functions {
			if strings.HasPrefix(funcName, objName+".") {
				suggestions = append(suggestions, funcName)
				if len(suggestions) >= 3 {
					break
				}
			}
		}
	} else {
		// Buscar funções de nível superior ou objetos com nomes similares
		for funcName := range functions {
			// Se for uma função de nível superior (sem ponto)
			if !strings.Contains(funcName, ".") {
				if len(name) > 0 && len(funcName) > 0 &&
					(strings.HasPrefix(funcName, name[:1]) ||
						levenshteinDistance(funcName, name) <= 3) {
					suggestions = append(suggestions, funcName)
				}
			} else if !strings.Contains(name, ".") {
				// Se estamos procurando por um nome que pode ser um objeto
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

// Função auxiliar para verificar se o tipo passado é válido para o tipo esperado
func isValidJSType(passedArg string, expectedType string) bool {
	// Como JavaScript é tipado dinamicamente, podemos ser flexíveis
	if expectedType == "any" {
		return true
	}

	// Para números
	if expectedType == "number" {
		// Verificar se é um número (não contém aspas e não é boolean)
		if strings.Contains(passedArg, "\"") || strings.Contains(passedArg, "'") {
			return false // É uma string
		}
		if passedArg == "true" || passedArg == "false" {
			return false // É um boolean
		}
	}

	// Para strings
	if expectedType == "string" {
		return strings.HasPrefix(passedArg, "\"") || strings.HasPrefix(passedArg, "'")
	}

	// Para boolean
	if expectedType == "boolean" {
		return passedArg == "true" || passedArg == "false"
	}

	// Aceitar por padrão (JavaScript é flexível)
	return true
}

// Implementação de Levenshtein distance para encontrar nomes similares
func levenshteinDistance(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	// Inicializar matriz de distâncias
	matrix := make([][]int, len(a)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
		matrix[i][0] = i
	}
	for j := range matrix[0] {
		matrix[0][j] = j
	}

	// Calcular distâncias
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deleção
				matrix[i][j-1]+1,      // inserção
				matrix[i-1][j-1]+cost, // substituição
			)
		}
	}

	return matrix[len(a)][len(b)]
}

// Função auxiliar para encontrar o valor mínimo entre três inteiros
func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

func (v *R2D2Visitor) isAccessibleFunction(funcName string) (bool, Function, string) {
	// Verificar se a função existe no módulo global
	globalModule, exists := v.symbolTable.Modules["global"]
	if !exists {
		return false, Function{}, "/* ERROR: Global module not initialized */"
	}

	// Verificar se a função existe no módulo atual (se estiver definido)
	// Precisamos verificar se o currentModule existe de outra forma, já que não é um ponteiro
	// Uma maneira de verificar é verificando se o campo Name não está vazio
	if v.currentModule.Name != "" {
		// Verificar se a função está no módulo atual
		function, exists := v.currentModule.Functions[funcName]
		if exists {
			return true, function, ""
		}
	}

	// Verificar no módulo global
	function, exists := globalModule.Functions[funcName]
	if exists {
		return true, function, ""
	}

	// Tente encontrar o objeto primeiro (ex: para 'console.log', buscamos 'console')
	objName := strings.Split(funcName, ".")[0]
	if strings.Contains(funcName, ".") {
		// Vamos ver se existem métodos para este objeto
		var objectMethods []string
		for fname := range globalModule.Functions {
			if strings.HasPrefix(fname, objName+".") {
				objectMethods = append(objectMethods, fname)
			}
		}

		if len(objectMethods) > 0 {
			// Se encontrarmos métodos para este objeto, sugerimos eles
			suggestions := objectMethods
			if len(suggestions) > 3 {
				suggestions = suggestions[:3]
			}

			errorMessage := fmt.Sprintf(
				"/* ERROR: Function '%s' not found. Você quis dizer: %s? */",
				funcName, strings.Join(suggestions, ", "),
			)
			return false, Function{}, errorMessage
		}
	}

	// Se não encontrarmos o objeto ou não for uma referência a objeto,
	// buscamos funções com nomes similares
	suggestions := findSimilarFunctions(globalModule.Functions, funcName)

	errorMessage := fmt.Sprintf("/* ERROR: Function '%s' not found */", funcName)
	if len(suggestions) > 0 {
		errorMessage += fmt.Sprintf(" /* Você quis dizer: %s? */", strings.Join(suggestions, ", "))
	}

	return false, Function{}, errorMessage
}
