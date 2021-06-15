package main

import (
	"api-rest-go-mysql/api"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// init inicializa cria o banco de dados e a tabela usuarios
func init() {
	db, err := sql.Open("mysql", "user:123456@/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	api.Execute(db, "create database if not exists databasego")
	api.Execute(db, "use databasego")
	api.Execute(db, "drop table if exists usuarios")
	api.Execute(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(100),
		email varchar (80),
		PRIMARY KEY (id)
	)`)

	// Inicializa a tabela "usuarios" com alguns registros
	stmt, _ := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	stmt.Exec("Maria Silva", "maria@email.com")
	stmt.Exec("João Almeida", "joao@email.com")
	stmt.Exec("Ricardo José", "ricjose@email.com")
}

func main() {
	fmt.Println("Servidor está rodando na porta 8080...")

	http.HandleFunc("/api/usuarios", api.ListarUsuarios)
	//endpoint: http://localhost:8080/api/usuarios (GET)
	http.HandleFunc("/api/usuarios/selecione", api.SelecionarUsuarios)
	//endpoint: http://localhost:8080/api/usuarios/selecione?id=1 (GET)
	http.HandleFunc("/api/usuarios/cadastro", api.CadastrarUsuario)
	//endpoint: http://localhost:8080/api/usuarios/cadastro (POST)
	http.HandleFunc("/api/usuarios/edicao", api.EditarUsuario)
	//endpoint: http://localhost:8080/api/usuarios/edicao (PUT)
	http.HandleFunc("/api/usuarios/delecao", api.DeletarUsuario)
	//endpoint: http://localhost:8080/api/usuarios/delecao?id=1 (DELETE)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
