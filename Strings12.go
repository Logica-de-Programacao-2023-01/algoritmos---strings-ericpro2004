package main

import (
	"fmt"
	"strings"
)

func main() {
	//Código de verificação de palídromo em string.
	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	if palindromo(nomes) {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}

func palindromo(nome string) bool {
	nome = strings.ToLower(strings.ReplaceAll(nome, " ", ""))

	tamanho := len(nome)
	for i := 0; i < tamanho/2; i++ {
		if nome[i] != nome[tamanho-1-i] {
			return false
		}
	}

	return true
}
