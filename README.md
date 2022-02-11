# api-rest-go-mysql
Simples API REST feita em Golang usando a biblioteca padrão, banco de dados MySql e execução no Docker

# Pre-requisitos:

	1 - docker e docker-compose;
	
# Executar a aplicação:

	// no terminal, vá até a raiz do projeto:
	1 - Para "buildar" o projeto na primeira execução:
	docker-compose build

	2 - Para executar a aplicação:
	docker-compose up

	3 - Aguarde o servidor mysql apresentar a mensagem "ready for connections" (control-c para finalizar)
	
# Endpoints - Para testar os serviços recomedo usar o Postman:

	// Para listar todos os usuários, executar:
	// http://localhost:8080/api/usuarios (GET)
	
	// Para selecionar usuário por id:
	// http://localhost:8080/api/usuarios/selecione?id=1 (GET)

	// Para cadastrar um novo usuário
	// Adicione o seguinte json no corpo da requisição, no formato "raw" e execute a requisição como exemplo abaixo:
	{
		"nome": "Thaís",
		"e-mail": "tata@email.com"
	}	
	// http://localhost:8080/api/usuarios/cadastro (POST)

	// Para editar um usuário existente
	// Adicione o seguinte json no corpo da requisição, no formato "raw" e execute a requisição como exemplo abaixo:
	{
		"id": 3,
		"nome": "nome alterado",
		"e-mail": "email_alterado@email.com"
	}
	// http://localhost:8080/api/usuarios/edicao (PUT)
	
	// Para deletar usuário por id, executar:
	// http://localhost:8080/api/usuarios/delecao?id=1 (DELETE)
