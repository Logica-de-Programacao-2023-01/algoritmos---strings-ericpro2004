package main

import (
	"fmt"
)

func main() {
	//Inversão de String.
	var texto string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&texto)

	stringinvertida := stringinvertida(texto)

	fmt.Println("String invertida:", stringinvertida)
}

func stringinvertida(input string) string {
	escrita := []rune(input)
	tamanho := len(escrita)

	for i := 0; i < tamanho/2; i++ {
		escrita[i], escrita[tamanho-1-i] = escrita[tamanho-1-i], escrita[i]
	}

	return string(escrita)
}
