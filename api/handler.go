package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

func codificarHeader(w http.ResponseWriter, codHttp int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(codHttp)
}

func codificarRetorno(w http.ResponseWriter, response interface{}, codHttp int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(codHttp)
	if !responseVazio(response) {
		json.NewEncoder(w).Encode(response)
	}
}

func responseVazio(response interface{}) (b bool) {
	b = false
	if reflect.ValueOf(response).IsZero() {
		b = true
	}
	return
}

func recuperarId(r *http.Request) (id string, err error) {
	id = r.URL.Query().Get("id")
	if id != "" {
		_, err = strconv.Atoi(id)
	}
	return id, err
}

func recuperarBody(r *http.Request) (usuario Usuario, err error) {
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &usuario)
	return usuario, err
}

// ListarUsuarios recebe uma requisição GET e retorna um usuário pelo Id ou todos os usuários
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		codificarHeader(w, http.StatusMethodNotAllowed)
		return
	}
	response, status := listarUsuariosRepo()
	codificarRetorno(w, response, status)
}

// SelecionarUsuarios recebe uma requisição GET e retorna a lista com todos os usuários
func SelecionarUsuarios(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		codificarHeader(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := recuperarId(r)
	if err != nil || id == "" {
		codificarHeader(w, http.StatusBadRequest)
		return
	}
	response, status := selecionarUsuarioRepo(id)
	codificarRetorno(w, response, status)
}

// CadastrarUsuario recebe uma requisição POST e cadastra um usuário enviado no corpo da requisição
func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		codificarHeader(w, http.StatusMethodNotAllowed)
		return
	}
	body, err := recuperarBody(r)
	if err != nil {
		codificarHeader(w, http.StatusBadRequest)
		return
	}
	status := cadastrarUsuarioRepo(body)
	codificarHeader(w, status)
}

// EditarUsuario recebe uma requisição PUT e edita um usuário enviado no corpo da requisição
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		codificarHeader(w, http.StatusMethodNotAllowed)
		return
	}
	body, err := recuperarBody(r)
	if err != nil {
		codificarHeader(w, http.StatusBadRequest)
		return
	}
	status := editarUsuarioRepo(body)
	codificarHeader(w, status)
}

// DeletarUsuario recebe uma requisição DELETE e apaga um usuário pelo Id
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		codificarHeader(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := recuperarId(r)
	if err != nil || id == "" {
		codificarHeader(w, http.StatusBadRequest)
		return
	}
	status := deletarUsuarioRepo(id)
	codificarHeader(w, status)
}
