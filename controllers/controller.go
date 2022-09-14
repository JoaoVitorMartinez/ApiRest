package controllers

import (
	"App/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func TodosOsJogos(w http.ResponseWriter, r *http.Request) {
	log.Println("Acesso em '/api/jogos'")

	json.NewEncoder(w).Encode(models.Jogos)

}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Acesso em '/'")

}

func BuscaJogoPorId(w http.ResponseWriter, r *http.Request) {
	log.Println("Acesso em '/api/jogos/{id}'")
	vars := mux.Vars(r)
	id := vars["id"]

	for _, jogo := range models.Jogos {
		if id == strconv.Itoa(jogo.Id) {
			json.NewEncoder(w).Encode(jogo)

		}
	}
}
