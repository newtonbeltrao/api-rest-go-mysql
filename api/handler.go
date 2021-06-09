package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func codificarResposta(w http.ResponseWriter, response interface{}, codHttp int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(codHttp)
	json.NewEncoder(w).Encode(response)
}

func codificarNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func recuperarBody(r *http.Request) Usuario {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var usuario Usuario
	json.Unmarshal(body, &usuario)
	return usuario
}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		var response interface{}
		if id != "" {
			response = selecionarUsuarioRepo(id)
		} else {
			response = listarUsuariosRepo()
		}
		codificarResposta(w, response, http.StatusOK)
	} else {
		codificarNotFound(w)
	}
}

func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		response := cadastrarUsuarioRepo(recuperarBody(r))
		codificarResposta(w, response, http.StatusCreated)
	} else {
		codificarNotFound(w)
	}
}

func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		response := editarUsuarioRepo(recuperarBody(r))
		codificarResposta(w, response, http.StatusAccepted)
	} else {
		codificarNotFound(w)
	}
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if r.Method == "DELETE" && id != "" {
		response := deletarUsuarioRepo(id)
		codificarResposta(w, response, http.StatusOK)
	} else {
		codificarNotFound(w)
	}
}
