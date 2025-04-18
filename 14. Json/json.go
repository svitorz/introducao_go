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
	/**
	*	Função Marshal
	* */
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

	/**
	*	Função Unmarshal
	* */

	newJsonDog := []byte(`{
			"nome": "Rex",
			"raça": "Dálmata", 
			"idade": 3
		}`)

	var newC cachorro

	if err := json.Unmarshal(newJsonDog, &newC); err != nil {
		log.Fatal(err)
	}

	fmt.Println(newC)
}
