package main

import (
	"errors"
	"fmt"
)

func main() {
	/**
		* Tipos de numeros inteiros.
		* Diferenciados pelo numero de bits.
		*
		* int baseado na arquitetura do computador.
		* int8, int16, int32, int64; // baseada na quantidade de bits
	* */
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint = 128 // unsygned int (int sem sinal)
	fmt.Println(numero2)

	/**
	* int32 = rune
	* uint8 = byte
	* */

	var realN float32 = 123.12
	var realN2 float64 = 1421241.12

	fmt.Println(realN, realN2)

	// não existe char
	char := 'A'
	fmt.Println(char)

	var sim bool = true
	var nao bool = false
	if sim != nao {
		fmt.Println(sim)
	}

	var err error = errors.New("Erro interno.") // valor default <nil>
	fmt.Println(err)
}
