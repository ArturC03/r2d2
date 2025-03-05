package visitor

//
// import (
// 	"fmt"
// 	r2d2Styles "github.com/ArturC03/r2d2Styles"
// 	"strconv"
// )
//
// // Função para verificar se a variável tem um valor válido
// func (v Variable) hasValidValue() bool {
// 	switch v.Type {
// 	case "i8", "i16", "i32", "i64":
// 		return isValidSignedInteger(v.Value)
// 	case "u8", "u16", "u32", "u64":
// 		return isValidUnsignedInteger(v.Value)
// 	case "f32", "f64":
// 		return isValidFloat(v.Value)
// 	case "bool":
// 		return isValidBool(v.Value)
// 	case "string":
// 		return true // Strings são sempre válidas
// 	default:
// 		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Invalid type %s for variable %s", v.Type, v.Name)))
// 		return false
// 	}
// }
//
// // Função para validar inteiros assinados
// func isValidSignedInteger(value string) bool {
// 	_, err := strconv.ParseInt(value, 10, 64)
// 	return err == nil
// }
//
// // Função para validar inteiros não assinados
// func isValidUnsignedInteger(value string) bool {
// 	num, err := strconv.ParseUint(value, 10, 64)
// 	return err == nil && num >= 0
// }
//
// // Função para validar floats
// func isValidFloat(value string) bool {
// 	_, err := strconv.ParseFloat(value, 64)
// 	return err == nil
// }
//
// // Função para validar booleanos
// func isValidBool(value string) bool {
// 	return value == "true" || value == "false"
// }
