package main

import "fmt"

func main() {
	// neste caso, o canal pode receber dois valores, pois está explicitado a capacidade dele.
	canal := make(chan string, 2)

	canal <- "Olá, mundo"
	canal <- "Programando em Go"
	mensagem := <-canal
	fmt.Println(mensagem)
}
