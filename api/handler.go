package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

// codificarHeader prepara o header da requisição conforme o código e o método HTTP
func codificarHeader(w http.ResponseWriter, r *http.Request, codHttp int) {
	if (codHttp == 200 || codHttp == 204) && r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(codHttp)
	} else {
		w.WriteHeader(codHttp)
		w.Write([]byte(fmt.Sprintf("%d %s\n", codHttp, http.StatusText(codHttp))))
	}
}

// codificarRetorno codifica a resposta e o header para retornar ao browser
func codificarRetorno(w http.ResponseWriter, r *http.Request, response interface{}, erro error) {
	if erro != nil {
		codificarHeader(w, r, http.StatusInternalServerError)
	} else if reflect.ValueOf(response).IsZero() {
		codificarHeader(w, r, http.StatusNoContent)
	} else if r.Method == http.MethodPost {
		codificarHeader(w, r, http.StatusCreated)
	} else if r.Method == http.MethodPut || r.Method == http.MethodDelete {
		codificarHeader(w, r, http.StatusOK)
	} else if r.Method == http.MethodGet {
		codificarHeader(w, r, http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

// recuperarId retorna se o id é um número válido
func recuperarId(r *http.Request) (id string, err error) {
	id = r.URL.Query().Get("id")
	if id != "" {
		_, err = strconv.Atoi(id)
	}
	return id, err
}

// recuperarBody retorna se o corpo da requisição é válida
func recuperarBody(r *http.Request) (usuario Usuario, err error) {
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &usuario)
	return usuario, err
}

// ListarUsuarios recebe uma requisição GET e retorna todos os usuários
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		codificarHeader(w, r, http.StatusMethodNotAllowed)
		return
	}
	response, erro := listarUsuariosRepo()
	codificarRetorno(w, r, response, erro)
}

// SelecionarUsuarios recebe uma requisição GET e retorna um usuário pelo Id
func SelecionarUsuarios(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		codificarHeader(w, r, http.StatusMethodNotAllowed)
		return
	}
	id, err := recuperarId(r)
	if err != nil || id == "" {
		codificarHeader(w, r, http.StatusBadRequest)
		return
	}
	response, erro := selecionarUsuarioRepo(id)
	codificarRetorno(w, r, response, erro)
}

// CadastrarUsuario recebe uma requisição POST e cadastra um usuário enviado no corpo da requisição
func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		codificarHeader(w, r, http.StatusMethodNotAllowed)
		return
	}
	body, err := recuperarBody(r)
	if err != nil {
		codificarHeader(w, r, http.StatusBadRequest)
		return
	}
	response, erro := cadastrarUsuarioRepo(body)
	codificarRetorno(w, r, response, erro)
}

// EditarUsuario recebe uma requisição PUT e edita um usuário enviado no corpo da requisição
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		codificarHeader(w, r, http.StatusMethodNotAllowed)
		return
	}
	body, err := recuperarBody(r)
	if err != nil {
		codificarHeader(w, r, http.StatusBadRequest)
		return
	}
	response, erro := editarUsuarioRepo(body)
	codificarRetorno(w, r, response, erro)
}

// DeletarUsuario recebe uma requisição DELETE e apaga um usuário pelo Id
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		codificarHeader(w, r, http.StatusMethodNotAllowed)
		return
	}
	id, err := recuperarId(r)
	if err != nil || id == "" {
		codificarHeader(w, r, http.StatusBadRequest)
		return
	}
	response, erro := deletarUsuarioRepo(id)
	codificarRetorno(w, r, response, erro)
}
