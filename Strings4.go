package main

import "fmt"

func main() {
	//Comparação de strings.
	var string1, string2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&string1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&string2)

	if string1 == string2 {
		fmt.Println("As strings são iguais.")
	} else {
		fmt.Println("As strings são diferentes.")
	}
}
