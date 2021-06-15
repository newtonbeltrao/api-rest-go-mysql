package api

import (
	"database/sql"
	"log"
)

// conectarBanco abre e retorna uma conexão com o banco de dados
func conectarBanco() (db *sql.DB) {
	db, err := sql.Open("mysql", "user:123456@/databasego")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// execute executa uma query e retorna um Result
func Execute(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// listarUsuariosRepo retorna a lista com todos os usuários e um erro
func listarUsuariosRepo() (retorno []Usuario, erro error) {
	db := conectarBanco()
	defer db.Close()
	rows, erro := db.Query("select id, nome, email from usuarios")
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()
	var u Usuario
	for rows.Next() {
		rows.Scan(&u.Id, &u.Nome, &u.Email)
		retorno = append(retorno, u)
	}
	return retorno, erro
}

// selecionarUsuarioRepo retorna o usuário selecionado e um erro
func selecionarUsuarioRepo(id string) (response Usuario, erro error) {
	db := conectarBanco()
	defer db.Close()
	erro = db.QueryRow("select id, nome, email from usuarios  where id = ?", id).Scan(&response.Id, &response.Nome, &response.Email)
	if erro == sql.ErrNoRows {
		return response, nil
	}
	return response, erro
}

// cadastrarUsuarioRepo insere um usuário e retorna a quantidade de registros inseridos e um erro
func cadastrarUsuarioRepo(usuario Usuario) (response int64, erro error) {
	db := conectarBanco()
	defer db.Close()
	stmt, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		return 0, err
	}
	response, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return response, erro
}

// editarUsuarioRepo edita o registro de um usuário e retorna a quantidade de registros editados e um erro
func editarUsuarioRepo(usuario Usuario) (response int64, erro error) {
	db := conectarBanco()
	defer db.Close()
	stmt, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(usuario.Nome, usuario.Email, usuario.Id)
	if err != nil {
		return 0, err
	}
	response, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return response, erro
}

// deletarUsuarioRepo deleta o usuário selecionado e retorna a quantidade de registros deletados e um erro
func deletarUsuarioRepo(id string) (response int64, erro error) {
	db := conectarBanco()
	defer db.Close()
	stmt, err := db.Prepare("delete from usuarios  where id = ?")
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	response, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return response, erro
}
