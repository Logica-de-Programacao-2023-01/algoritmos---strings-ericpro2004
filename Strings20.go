package main

import (
	"fmt"
	"strings"
)

func main() {
	//Verificação de string para ver se ela está inserida no camelcase.
	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	isCamelCase := checkCamelCase(nomes)
	numplavras := countWords(nomes)

	if isCamelCase {
		fmt.Println("A string está em camelCase.")
	} else {
		fmt.Println("A string não está em camelCase.")
	}

	fmt.Println("Número de palavras:", numplavras)
}

func checkCamelCase(input string) bool {
	if len(input) == 0 {
		return false
	}

	firstChar := input[0]
	if !maiuscula(firstChar) {
		return false
	}

	for i := 1; i < len(input); i++ {
		char := input[i]
		if maiuscula(char) {
			return false
		}
	}

	return true
}

func maiuscula(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

func countWords(input string) int {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return !maiuscula(byte(r))
	})
	return len(words)
}
