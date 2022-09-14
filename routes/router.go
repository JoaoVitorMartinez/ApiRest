package routes

import (
	"App/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()

	fmt.Println("Servidor inicializado...")

	r.HandleFunc("/", controllers.Home).Methods("GET")

	r.HandleFunc("/api/jogos", controllers.TodosOsJogos).Methods("GET")

	r.HandleFunc("/api/jogos/{id}", controllers.BuscaJogoPorId).Methods("GET")

	http.ListenAndServe(":8000", r)
}
