package main

import (
	"fmt"
	"strings"
)

func main() {
	//Substituição de vogais de uma string por *.
	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	Nstring := VR(nomes)

	fmt.Println("String com vogais substituídas:", Nstring)
}

func VR(input string) string {
	vogais := "aeiouAEIOU"
	Repositor := strings.NewReplacer(vogais, "*")

	return Repositor.Replace(input)
}
