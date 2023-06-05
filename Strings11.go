package main

import (
	"fmt"
	"strings"
)

func main() {
	//Remoção de vogais de uma string.
	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	vogaisremovidas := vogaisremovidas(nomes)

	fmt.Println("String sem vogais:", vogaisremovidas)
}

func vogaisremovidas(input string) string {
	vogais := "aeiouAEIOU"
	resultado := ""

	for _, letra := range input {
		if !strings.ContainsRune(vogais, letra) {
			resultado += string(letra)
		}
	}
	return resultado
}
