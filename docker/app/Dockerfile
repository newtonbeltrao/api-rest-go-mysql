FROM golang
# Expõe a porta 8080
EXPOSE 8080
# Cria o diretório /build
RUN mkdir /build
# Adiciona os arquivos do diretório local para o diretório no docker
ADD . /build/
# Define o "workdir"
WORKDIR /build
# Build da aplicação
RUN go build -o ./build/main
# Executa a aplicação
CMD [ "./build/main" ]