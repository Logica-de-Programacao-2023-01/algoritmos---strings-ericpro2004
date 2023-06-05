package main

import (
	"fmt"
	"strings"
)

func main() {
	//Conversão de nome solicitado de letras minúsculas para maiúsculas.
	var nome string
	fmt.Print("Digite um nome: ")
	fmt.Scanln(&nome)

	LM := strings.ToUpper(nome)

	fmt.Println("String convertida para maiúsculas:", LM)
}
