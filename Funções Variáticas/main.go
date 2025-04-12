package main

import "fmt"

func main() {
	variativa("Olá, mundo", 1, 2, 4, 5, 6, 9, 7, 8)
}

func variativa(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}
