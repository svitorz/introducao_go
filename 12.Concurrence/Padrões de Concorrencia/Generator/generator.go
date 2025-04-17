package main

import (
	"fmt"
	"time"
)

func escrever(texto string) <-chan string {
	canal := make(chan string)

	// a função encapsula a goroutine
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	// e retorna o canal
	return canal
}

func main() {
	canal := escrever("Olá, mundo!")
	for i := 0; i <= 10; i++ {
		fmt.Println(<-canal)
	}
}
