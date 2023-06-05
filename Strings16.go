package main

import (
	"fmt"
	"strings"
)

func main() {
	//Verificação de duas strings (se uma é substring da outra).
	var str1, str2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	if Substring(str1, str2) {
		fmt.Println("A segunda string é uma substring da primeira.")
	} else {
		fmt.Println("A segunda string não é uma substring da primeira.")
	}
}

func Substring(str1, str2 string) bool {
	return strings.Contains(str1, str2)
}
