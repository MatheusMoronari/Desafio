package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MatheusMoronari/Desafio/entidades"
	"github.com/MatheusMoronari/Desafio/models"
	_ "github.com/lib/pq"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var dadoslogin entidades.Dadoslogin
	error := json.NewDecoder(r.Body).Decode(&dadoslogin)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}
	if models.ValidaUsuario(dadoslogin) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Login bem-sucedido"})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Usuário ou senha inválidos"})
	}
}
