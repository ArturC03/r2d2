package visitor

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

var availableFunctions map[string][]string

// Carregar as funções globais do JavaScript usando Deno
func loadGlobalFunctions() error {
	// Comando Deno para obter as funções globais
	cmd := exec.Command("deno", "eval", `
const availableFunctions = {};

Object.getOwnPropertyNames(globalThis).forEach(name => {
  try {
    const obj = globalThis[name];
    if (typeof obj === 'object' || typeof obj === 'function') {
      availableFunctions[name] = Object.getOwnPropertyNames(obj)
        .filter(prop => typeof obj[prop] === 'function');
    }
  } catch (e) {
    availableFunctions[name] = ['<inaccessible>'];
  }
});

console.log(JSON.stringify(availableFunctions, null, 2));
`)

	// Capturar a saída do comando
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao executar o comando Deno: %v", err)
	}

	// Fazer o parse da saída JSON para a variável global availableFunctions
	if err := json.Unmarshal(output, &availableFunctions); err != nil {
		return fmt.Errorf("erro ao parsear o JSON: %v", err)
	}

	return nil
}
