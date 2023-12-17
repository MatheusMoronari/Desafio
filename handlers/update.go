package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/MatheusMoronari/Desafio/models"
	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer conversao do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var pessoa models.Pessoa

	err = json.NewDecoder(r.Body).Decode(&pessoa)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	rows, err := models.Update(int64(id), pessoa)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows > 1 {
		log.Printf("Error: foram atualizados %d", rows)
	}
	resp := map[string]any{
		"Error":   false,
		"Message": "dados atualizados com sucesso!",
	}
	w.Header().Add("Content=Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
