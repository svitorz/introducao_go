package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raça"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 4}
	fmt.Println(c)

	jsonDog, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	// printa em formato de bytes
	fmt.Println(jsonDog)

	// printa em formato Json
	fmt.Println(bytes.NewBuffer(jsonDog))
}
