package controllers

import (
	"App/database"
	"App/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func TodosOsJogos(w http.ResponseWriter, r *http.Request) {
	log.Println("Acesso em '/api/jogos'")
	var j []models.Jogo
	database.DB.Find(&j)

	json.NewEncoder(w).Encode(j)

}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Acesso em '/'")

}

func BuscaJogoPorId(w http.ResponseWriter, r *http.Request) {
	log.Println("Acesso em '/api/jogos/{id}'")
	vars := mux.Vars(r)
	id := vars["id"]
	var j []models.Jogo
	database.DB.First(&j, id)

	json.NewEncoder(w).Encode(j)
}

func InserirJogo(w http.ResponseWriter, r *http.Request) {
	var novoJogo models.Jogo

	json.NewDecoder(r.Body).Decode(&novoJogo)
	database.DB.Create(&novoJogo)
	json.NewEncoder(w).Encode(novoJogo)
}
