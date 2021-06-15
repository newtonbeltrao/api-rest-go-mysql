# api-rest-go-mysql
Simples API REST feita em Golang usando a biblioteca padrão e banco de dados MySql

# Pre-requisitos:

1 - Instalar Mysql e criar um usuário "user", senha "123456" e adicionar o grant ao banco de dados "databasego";

	// conectado ao mysql:
	create user 'user'@'localhost' identified by '123456';
	grant all on databasego.* to 'user'@'localhost';

2 - Instalar o driver mysql;

	// no terminal do projeto Go:
	go get -u github.com/go-sql-driver/mysql
	
# Executar a aplicação:

	// no terminal do projeto Go:
	go run main.go
	
# Endpoints - Para testar os serviços recomedo usar o Postman:

	// Listar todos os usuários
	// http://localhost:8080/api/usuarios (GET)
	
	// Selecionar usuário por id
	// http://localhost:8080/api/usuarios/selecione?id=1 (GET)

	// Cadastrar um novo usuário
	// http://localhost:8080/api/usuarios/cadastro (POST)

	// Editar um usuário existente
	// http://localhost:8080/api/usuarios/edicao (PUT)
	
	// Deletar usuário por id
	// http://localhost:8080/api/usuarios/delecao?id=1 (DELETE)
