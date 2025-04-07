package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	// var u usuario
	// u.nome = "Vitor"
	// u.idade = 18
	//
	// caso só tenha um dos valores, pode ser definido como
	// usuario := usuario{nome: "Vitor"}
	// e o outro valor pode continuar nulo.
	u := usuario{"Vitor", 18}
	fmt.Println(u)
}
