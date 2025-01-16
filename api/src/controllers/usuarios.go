package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal("erro no corpo request", erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal("erro no unmarshal", erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal("erro em conectar com o banco", erro)
	}

	repositorio := repositorio.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)
	usuarioId, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal("erro em criar usuario", erro)
	}

	w.Write([]byte(fmt.Sprintf("Usuario inserido %d", usuarioId)))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuario"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuarios"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario"))
}
