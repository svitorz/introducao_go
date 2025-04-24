package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "root:mydba@tcp(127.0.0.1:3307)/introducao_go?charset=utf8&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
