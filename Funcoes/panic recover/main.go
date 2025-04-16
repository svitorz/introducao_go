package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentando recuperar a execução.")
}

func alunoEstaAprovado(n1, n2 int) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	// quebra a execução do sistema, chamando todos os defer.
	panic("A média é exatamente 6.")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 7))
	fmt.Println("Pós execução")
}
