package main

import "fmt"

func main() {
	array1 := [5]string{"v", "i", "t", "o", "r"}
	fmt.Println(array1)

	// slices
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}
	fmt.Println(slice)

	// Arrays internos
	slice2 := make([]float32, 10, 15)
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // tamanho
	fmt.Println(cap(slice2)) // capacidade
	/**
	* Se o tamanho do array interno "estourar", o compilador irá duplicá-lo, consequentemente, dobrando de tamanho.
	* */
	//podemos resumir que o array tem um tamanho fixo, e o slice não precisa ter.
	//Já o array, pode "dobrar" de tamanho de acordo com a necessidade.
}
