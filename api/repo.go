package api

import (
	"database/sql"
	"log"
	"reflect"
)

// execute executa uma query e retorna um Result
func Execute(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// listarUsuariosRepo retorna a lista com todos os usuários
func listarUsuariosRepo() (retorno []Usuario, status int) {
	db, err := sql.Open("mysql", "user:123456@/databasego")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select id, nome, email from usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var u Usuario
	for rows.Next() {
		rows.Scan(&u.Id, &u.Nome, &u.Email)
		retorno = append(retorno, u)
	}
	if reflect.ValueOf(retorno).IsZero() {
		status = 404
	} else {
		status = 200
	}
	return retorno, status
}

// selecionarUsuarioRepo retorna o usuário selecionado
func selecionarUsuarioRepo(id string) (retorno Usuario, status int) {
	db, err := sql.Open("mysql", "user:123456@/databasego")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.QueryRow("select id, nome, email from usuarios  where id = ?", id).
		Scan(&retorno.Id, &retorno.Nome, &retorno.Email)
	if reflect.ValueOf(retorno).IsZero() {
		status = 404
	} else {
		status = 200
	}
	return retorno, status
}

// cadastrarUsuarioRepo insere um novo usuário e retorna a nova lista de usuários
func cadastrarUsuarioRepo(usuario Usuario) (status int) {
	db, err := sql.Open("mysql", "user:123456@/databasego")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	result, err := stmt.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		log.Fatal(err)
	}
	rows, _ := result.RowsAffected()
	if rows == 1 {
		status = 201
	} else {
		status = 400
	}
	return status
}

// editarUsuarioRepo edita o registro de um usuário e retorna a lista de usuários
func editarUsuarioRepo(usuario Usuario) (status int) {
	db, err := sql.Open("mysql", "user:123456@/databasego")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	result, err := stmt.Exec(usuario.Nome, usuario.Email, usuario.Id)
	if err != nil {
		log.Fatal(err)
	}
	rows, _ := result.RowsAffected()
	if rows == 1 {
		status = 201
	} else {
		status = 404
	}
	return status
}

// deletarUsuarioRepo deleta o usuário selecionado e retorna a lista de usuários restantes
func deletarUsuarioRepo(id string) (status int) {
	db, err := sql.Open("mysql", "user:123456@/databasego")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("delete from usuarios  where id = ?")
	result, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	rows, _ := result.RowsAffected()
	if rows == 1 {
		status = 200
	} else {
		status = 404
	}
	return status
}
