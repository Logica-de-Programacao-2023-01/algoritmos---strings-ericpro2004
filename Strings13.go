package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Código para verificar se a string solicitada é uma sequência númerica crescente.
	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	if sequencia(nomes) {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}

func sequencia(input string) bool {
	tamanho := len(input)
	if tamanho < 2 {
		return false
	}

	for i := 1; i < tamanho; i++ {
		atual, err1 := strconv.Atoi(string(input[i]))
		anterior, err2 := strconv.Atoi(string(input[i-1]))

		if err1 != nil || err2 != nil || atual <= anterior {
			return false
		}
	}

	return true
}
