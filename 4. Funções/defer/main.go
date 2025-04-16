package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função função 2")
}

func calculaMedia(n1, n2 int) bool {
	defer fmt.Println("Media calculada. Resultado será retornado")
	fmt.Println("Entrando na funçao para verificar se o aluno esta aprovado")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao1()
	// defer = adiar
	funcao2()

	fmt.Println(calculaMedia(10, 7))
}
