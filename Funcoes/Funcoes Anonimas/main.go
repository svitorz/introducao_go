package main

import "fmt"

func main() {
	retorno := func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("Isto é uma função anonima.")

	fmt.Println(retorno)
}
