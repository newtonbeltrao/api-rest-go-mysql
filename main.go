package main

import (
	"api-rest-go-mysql/api"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

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
