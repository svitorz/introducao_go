package main

import "fmt"

func main() {
	// tipagem map[chave]valor
	usuario := map[string]string{
		"nome":      "Vitor",
		"sobrenome": "Souza",
	}

	fmt.Println(usuario)
	delete(usuario, "nome")
	fmt.Println(usuario)

	usuario["email"] = "vitor@test.com"
	fmt.Println(usuario)
}
