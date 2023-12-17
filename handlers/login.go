package handlers

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Obter usuário e senha dos parâmetros da requisição
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	// Validar usuário
	if validateUser(username, password) {
		fmt.Fprintf(w, "Login bem-sucedido")
	} else {
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
	}
}
