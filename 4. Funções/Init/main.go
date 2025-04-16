package main

import "fmt"

var n int

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
	// se o valor de N nao for iniciado na função init, ele sera printado como 0.
}

/**
* A função init é executada antes de qualquer outra função
* */
func init() {
	n = 10
	// agora, o print na func main, mostra n como "10"
	fmt.Println("Executando a função init")
}
