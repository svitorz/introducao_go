package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá, Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")
	for {
		mensagem, aberto := <-canal

		if !aberto {
			break
		}

		fmt.Println(mensagem)
	}
}

func escrever(text string, canal chan string) {
	// time.Sleep(time.Second * 5)
	for i := 0; i <= 5; i++ {
		canal <- text
		time.Sleep(time.Second)
	}

	// fecha o canal, impedindo uma deadlock
	close(canal)
}
