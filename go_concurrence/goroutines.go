package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// utilizando waitgroups
	var waitGroup sync.WaitGroup

	// rodando duas goroutines por vez
	waitGroup.Add(2)

	go func() {
		escrever("Olá, mundo!")
		waitGroup.Done() //-1 goroutine
	}()

	// O sistema nao espera a finalização de um processo para iniciar outro.
	// Neste caso abaixo, nao esta implementada a waitGroup
	// go escrever("Olá, mundo!")

	go func() {
		escrever("Programando em GO!")
		waitGroup.Done() //-1 goroutine
	}()

	waitGroup.Wait() // 0
}

func escrever(texto string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
