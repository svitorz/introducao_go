package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		w.Write([]byte("Erro ao converter usuário para struc"))
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

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		if err != nil {
			w.Write([]byte("Erro ao converter ID do usuario"))
			return
		}
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	row, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao realizar operação"))
		return
	}

	var usuarios []usuario

	if row.Next() {
		var usuario usuario
		if err := row.Scan(&usuario.ID, &usuario.Email, &usuario.Nome); err != nil {
			w.Write([]byte("Erro ao escanear usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter usuarios em Json"))
		return
	}
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	rows, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao realizar operação"))
		return
	}

	var usuarios []usuario

	for rows.Next() {
		var usuario usuario
		if err := rows.Scan(&usuario.ID, &usuario.Email, &usuario.Nome); err != nil {
			w.Write([]byte("Erro ao escanear usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter usuarios em Json"))
		return
	}
}
