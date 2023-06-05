package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Verificação de números 0 a 9 em strings.
	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	if CN(nomes) {
		fmt.Println("A string contém somente números.")
	} else {
		fmt.Println("A string não contém somente números.")
	}
}

func CN(nomes string) bool {
	_, err := strconv.Atoi(nomes)
	return err == nil
}
