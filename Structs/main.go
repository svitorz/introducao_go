package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	sobrenome string
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

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
	//u := usuario{"Vitor", 18}
	// fmt.Println(u)

	p := pessoa{"Vitor", 18, "Souza", 175}
	fmt.Println(p)
	e := estudante{p, "Sistemas de Informação", "IFSP"}
	fmt.Println(e)
}
