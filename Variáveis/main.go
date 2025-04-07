package main

import "fmt"

func main() {
	// tipagem explícita
	var name string = "Vitor"

	/***
	* tipagem implícita
	* inferência de tipo
	 */

	another_name := "Vitor Fábio"
	fmt.Println(name)

	fmt.Println(another_name)

	const constName string = "Vitor Contante"
	fmt.Println(constName)

	/**
	* Invertendo valores de variáveis.
	* */

	var (
		objeto string = "Controle"
		casa   string = "mesa"
	)
	fmt.Println(objeto, casa)
	objeto, casa = casa, objeto
	fmt.Println(objeto, casa)
}
