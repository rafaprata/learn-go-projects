package main

import (
	"fmt"
	"unicode"
	"math"
)

var inputValue string
var convertedValue int

func readValue() string {
	fmt.Print("Digite o valor em binario: ")
	fmt.Scan(&inputValue)
	return inputValue
}

func validateValue(value string) bool {
	// verificar se existem letras e se e binario	
	for _, char := range value {
		if unicode.IsLetter(char) {
			fmt.Print("[ERROR] The input is not a NUMBER\n")
			return false
		}
		if char != '0' && char != '1' {
			fmt.Print("[ERROR] The number is not BINARY\n")
			return false
		}
	}
	// verificar se tem tamanho 8
	if len(value) > 8 {
		fmt.Print("[ERROR] The number is longer than 8\n")
		return false
	}
	return true
}

func completeBinary(value string) string {
	if len(value) < 8 {
		complementValue := fmt.Sprintf("%08s", value) // Completa com zeros à esquerda até que tenha 8 caracteres
		return complementValue
	}
	return value
}

func convertValue(value string) {
	complementedInputValue := completeBinary(value)
	convertedValue = 0 // Resetar o valor convertido para garantir que não haja acumulação
	for i, char := range complementedInputValue {
		expoente := 7 - i
		if char == '1' {
			convertedValue += int(math.Pow(2, float64(expoente)))
		}
		
	}
}

func main() {
	inputValue := readValue()
	if !validateValue(inputValue) {
		return
	}
	convertValue(inputValue)
	fmt.Printf("%v\n", convertedValue)
}