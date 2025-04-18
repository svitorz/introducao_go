package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))

	u := usuario{"Vitor", "vitor@gmail.com"}
	templates.ExecuteTemplate(w, "index.html", u)
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", index)

	fmt.Println("Listen on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
