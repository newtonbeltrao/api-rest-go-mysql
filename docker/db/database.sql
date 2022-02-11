-- Cria e utiliza o banco de dados
CREATE DATABASE IF NOT EXISTS databasego CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE databasego;

-- Cria o usuário "user" e concede o grant para acessar o banco no serviço "db" do container
CREATE USER IF NOT EXISTS 'user'@'db' IDENTIFIED BY '123456';
GRANT ALL PRIVILEGES ON databasego.* TO 'user'@'db' WITH GRANT OPTION;
FLUSH PRIVILEGES;

-- Apaga e recria a tabela "usuários"
DROP TABLE IF EXISTS usuarios;
CREATE TABLE usuarios (
    id integer auto_increment,
    nome varchar(80),
	email varchar (80),
    PRIMARY KEY (id)
);

-- Inicializa a tabela com alguns registros
insert into usuarios (nome, email) values ("Maria", "maria@email.com");
insert into usuarios (nome, email) values ("José", "jose@email.com");
insert into usuarios (nome, email) values ("Paulo", "paulo@email.com");