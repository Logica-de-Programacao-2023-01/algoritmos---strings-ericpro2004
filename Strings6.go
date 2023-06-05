package main

import (
	"fmt"
	"strings"
)

func main() {
	//Número de palavras em string.
	var frase string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&frase)

	numpalavras := Contadorpalavras(frase)

	fmt.Println("A string contém", numpalavras, "palavras.")
}

func Contadorpalavras(frase string) int {
	Espaço := strings.TrimSpace(frase)
	palavras := strings.Fields(Espaço)
	return len(palavras)
}
