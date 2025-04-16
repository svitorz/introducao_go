package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções recursivas.")

	for i := uint(0); i <= 15; i++ {
		fmt.Println(fibonacci(i))
	}
}
