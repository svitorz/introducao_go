package main

import (
	"fmt"
	"math/rand"
	"time"
)

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	// a função encapsula a goroutine
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	// e retorna o canal
	return canal
}

func ain() {
	canal := multiplexar(escrever(("Olá, Mundo!")), escrever("Programando em Go."))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
