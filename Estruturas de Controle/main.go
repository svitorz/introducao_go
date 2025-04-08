package main

import "fmt"

func main() {
	if numero := 15; numero > 10 {
		fmt.Println("Maior que 15")
	}
	// numero não funciona fora desse escopo do if.
	fmt.Println(diaDaSemana(2))
}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo."
	case 2:
		return "Segunda-feira."
	case 3:
		return "Terça-feira."
	case 4:
		return "Quarta-feira."
	case 5:
		return "Quinta-feira."
	case 6:
		return "Sexta-feira."
	case 7:
		return "Sábado"
	default:
		return "Inválido."
	}
}
