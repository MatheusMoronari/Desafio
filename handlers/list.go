package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MatheusMoronari/Desafio/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	pessoas, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter Registros: %v", err)
	}
	w.Header().Add("Content=Type", "application/json")
	json.NewEncoder(w).Encode(pessoas)
}
