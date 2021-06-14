# api-rest-go-mysql
Simples API REST feita em Golang usando banco de dados MySql

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
	// http://localhost:8080/api/listar (GET)
	
	// Selecionar usuário por id
	// http://localhost:8080/api/selecionar?id=1 (GET)

	// Cadastrar um novo usuário
	// http://localhost:8080/api/cadastrar (POST)

	// Editar um usuário existente
	// http://localhost:8080/api/editar (PUT)
	
	// Deletar usuário por id
	// http://localhost:8080/api/deletar?id=1 (DELETE)
