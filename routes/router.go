package routes

import (
	"App/controllers"
	"App/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	log.Println("Servidor inicializado.")

	r.HandleFunc("/", controllers.Home).Methods("GET")

	r.HandleFunc("/api/jogos", controllers.TodosOsJogos).Methods("GET")

	r.HandleFunc("/api/jogos/{id}", controllers.BuscaJogoPorId).Methods("GET")

	r.HandleFunc("/api/inserir", controllers.InserirJogo).Methods("POST")

	r.HandleFunc("/api/delete/{id}", controllers.DeletaJogo).Methods("DELETE")

	r.HandleFunc("/api/atualizar/{id}", controllers.AtualizaJogo).Methods("PUT")

	http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))
}
