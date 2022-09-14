package controllers

import (
	"App/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func TodosOsJogos(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(models.Jogos)

}

func Home(w http.ResponseWriter, r *http.Request) {
	println("Não faz nada por enquanto...")

}

func BuscaJogoPorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, jogo := range models.Jogos {
		if id == strconv.Itoa(jogo.Id) {
			json.NewEncoder(w).Encode(jogo)

		}
	}
}
