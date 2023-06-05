package main

import (
	"fmt"
	"strings"
)

func main() {
	//Substituição de palavras solicitadas pelo usuário.
	var nome, Al, Nl string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&Al)

	fmt.Print("Digite o novo caractere: ")
	fmt.Scanln(&Nl)

	NS := strings.ReplaceAll(nome, Al, Nl)

	fmt.Println("String atualizada:", NS)
}
