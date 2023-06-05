package main

import (
	"fmt"
	"unicode"
)

func main() {
	//Verificação de números em string.
	var texto string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&texto)

	num := false

	for _, Letra := range texto {
		if unicode.IsDigit(Letra) {
			num = true
			break
		}
	}

	if num {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém números.")
	}
}
