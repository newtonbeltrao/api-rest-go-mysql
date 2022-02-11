# Derivando da imagem oficial do MySQL
FROM mysql:8.0
# Adicionando os scripts SQL para serem executados na criação do banco
COPY ./docker/db/database.sql /docker-entrypoint-initdb.d/
# Expõe a porta do mysql
EXPOSE 3306