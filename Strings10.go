package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//Código para verificar se uma string solicitada é um anagrama.
	var str1, str2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	if éanagra(str1, str2) {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

func éanagra(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1 = organizaçãodestring(str1)
	str2 = organizaçãodestring(str2)

	return str1 == str2
}

func organizaçãodestring(str string) string {
	letras := strings.Split(str, "")
	sort.Strings(letras)
	return strings.Join(letras, "")
}
