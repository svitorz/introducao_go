package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

/**
* Função utilizando um ponteiro.
* */
func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	/**
	* var numero int = 20
	* // neste caso, foi passada uma cópia da variavel numero.
	* var numeroInvertido int = inverteSinal(numero)
	* fmt.Println(numeroInvertido)
	 */
	novoNumero := 40
	fmt.Println(novoNumero)
	inverteSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
