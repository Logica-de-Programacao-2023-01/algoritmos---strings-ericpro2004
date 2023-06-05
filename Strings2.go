package main

import (
	"fmt"
	"strings"
)

func main() {
	//Remoção de espaços de uma string.
	var nome string
	fmt.Print("Digite um texto(nome, frase, etc): ")
	fmt.Scanln(&nome)

	Semespaço := strings.ReplaceAll(nome, " ", "")

	fmt.Println("String sem espaços em branco: ", Semespaço)
}
