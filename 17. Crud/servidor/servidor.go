package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}
	var usuario usuario
	if err = json.Unmarshal(request, &usuario); err != nil {
		w.Write([]byte("Erro ao converter usuário para struct"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?,?)")
	if err != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	insert, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Erro ao inserir no banco de dados"))
		return
	}

	idInserido, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter id inserido"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}
