package main

import "fmt"

type user struct {
	name string
	age  uint8
}

// Esta função esta vinculada a estrutura user, como um método.
func (u user) salvar() {
	fmt.Printf("Salvando usuário %s.", u.name)
}

func (u user) maiorDeIdade() bool {
	return u.age >= 18
}

func (u *user) fazerAniversario() {
	u.age++
}

func main() {
	usuario1 := user{"Vitor Fábio", 18}
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())
}
