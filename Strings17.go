package main

import (
	"fmt"
)

func main() {
	//Verificação de letras únicas de uma string.
	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	LU := getUniqueLetters(nomes)

	fmt.Println("Letras únicas na string:", LU)
}

func getUniqueLetters(input string) string {
	unicas := ""
	javistas := make(map[rune]int)

	for _, letras := range input {
		javistas[letras]++
	}

	for _, char := range input {
		if javistas[char] == 1 {
			unicas += string(char)
		}
	}

	return unicas
}
