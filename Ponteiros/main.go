package main

import "fmt"

func main() {
	var n1 int = 10
	n2 := n1
	fmt.Println(n1, n2)
	// saída 10 10
	n1++
	fmt.Println(n1, n2)
	// saída 11 10

	// um ponteiro é uma referência em memória
	var variavel int = 100
	var ponteiro *int = &variavel
	fmt.Println(variavel, *ponteiro) // com o asterísco ocorre a desferenciação
	// saída sem asterísco 100 0x0001024
	// saída com * 100 100
}
