package main

import "fmt"

func closure() func() {
	var texto string = "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
