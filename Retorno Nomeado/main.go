package main

import "fmt"

// exemplo de função com retorno nomeado.
func calculadora(a int, b int) (soma int, multiplicacao int) {
	soma = a + b
	multiplicacao = a * b
	return
}

// exemplo de função variática
func soma(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}

func main() {
	fmt.Println(calculadora(10, 20))
	fmt.Println(soma(1, 2, 2, 3, 4, 4, 5, 5, 6, 3, 7, 3, 5, 4))
}
