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
	var dia string

	switch numero {
	case 1:
		dia = "Domingo."
		fallthrough // faz com que ele pule o case. Neste caso, ele retorna o que está no default, porque todos tem a clausula fallthrough
	case 2:
		dia = "Segunda-feira."
		fallthrough
	case 3:
		dia = "Terça-feira."
		fallthrough
	case 4:
		dia = "Quarta-feira."
		fallthrough
	case 5:
		dia = "Quinta-feira."
		fallthrough
	case 6:
		dia = "Sexta-feira."
		fallthrough
	case 7:
		dia = "Sábado"
		fallthrough
	default:
		dia = "Inválido."
	}
	return dia
}
