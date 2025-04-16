package main

import "fmt"

func somar(a int8, b int8) int8 {
	return a + b
}

func main() {
	// soma := somar(10, 20)
	// fmt.Println(soma)

	f := func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Olá, mundo.")
	fmt.Println(resultado)

	resultadosSoma, resultadoSubtracao := calculosMatematicos(10, 8)
	fmt.Println(resultadosSoma, resultadoSubtracao)

	// para ignorar um dos retornos da função, utilizo:
	resultados1, _ := calculosMatematicos(10, 8)
	fmt.Println(resultados1)
}

// neste caso, a tipagem do segundo parâmetro serve para o primeiro.
func calculosMatematicos(n, m int8) (int8, int8) {
	soma := n + m
	subtracao := n - m
	// permite dois retornos
	return soma, subtracao
}
