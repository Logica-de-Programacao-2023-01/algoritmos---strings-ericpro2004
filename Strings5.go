package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Verificação de String em ponto flutuante.
	var nome string

	fmt.Print("Digite um número em ponto flutuante: ")
	fmt.Scanln(&nome)

	_, erro := strconv.ParseFloat(nome, 64)

	if erro == nil {
		fmt.Println("É um número válido em ponto flutuante.")
	} else {
		fmt.Println("Não é um número válido em ponto flutuante.")
	}
}
