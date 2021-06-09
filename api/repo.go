package api

import (
	"database/sql"
	"log"
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
func listarUsuariosRepo() (retorno []Usuario) {
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
	return retorno
}

// selecionarUsuarioRepo retorna o usuário selecionado
func selecionarUsuarioRepo(id string) (retorno Usuario) {
	db, err := sql.Open("mysql", "user:123456@/databasego")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.QueryRow("select id, nome, email from usuarios  where id = ?", id).
		Scan(&retorno.Id, &retorno.Nome, &retorno.Email)

	return retorno
}

// cadastrarUsuarioRepo insere um novo usuário e retorna a nova lista de usuários
func cadastrarUsuarioRepo(usuario Usuario) (retorno []Usuario) {
	db, err := sql.Open("mysql", "user:123456@/databasego")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("insert into usuarios (nome, email) values (?, ?)", usuario.Nome, usuario.Email)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	retorno = listarUsuariosRepo()
	return retorno
}

// editarUsuarioRepo edita o registro de um usuário e retorna a lista de usuários
func editarUsuarioRepo(usuario Usuario) (retorno []Usuario) {
	db, err := sql.Open("mysql", "user:123456@/databasego")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("update usuarios set nome = ?, email = ? where id = ?", usuario.Nome, usuario.Email, usuario.Id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	retorno = listarUsuariosRepo()
	return retorno
}

// deletarUsuarioRepo deleta o usuário selecionado e retorna a lista de usuários restantes
func deletarUsuarioRepo(id string) (retorno []Usuario) {
	db, err := sql.Open("mysql", "user:123456@/databasego")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("delete from usuarios  where id = ?", id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	retorno = listarUsuariosRepo()
	return retorno
}
