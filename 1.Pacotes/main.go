package main

/**
* go mod tidy remove todas as dependências não utilizadas.
* */
import (
	"fmt"
	"modulo/auxiliar"

	// para importar utilize go get <package name>
	"github.com/badoux/checkmail"
)

/**
* Registra uma mensagem na tela.
* @return void
* Para executar todo o módulo estou utilizandos[]
* go run main.go
* caso quisesse buildar utilizaria go build
* */
func main() {
	fmt.Println("Escrevendo do arquivo main.")
	// funções em arquivos de pacotes diferentes, devem estar com a primeira lestra maiúscula para que possam ser utilizados dentro do módulo.

	auxiliar.Escrever()

	err := checkmail.ValidateFormat("vitor@gmail.com")
	fmt.Println(err)
}
