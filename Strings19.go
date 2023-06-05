package main

import "fmt"

func main() {
	////Inversão de String.
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	reversedString := reverseString(input)

	fmt.Println("String invertida:", reversedString)
}

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
