package main

import (
	"App/database"
	"App/models"
	"App/routes"
	"log"
)

func main() {
	log.Println("inicializando Servidor...")

	models.Jogos = []models.Jogo{
		{Id: 1, Nome: "Far Cry 6", Descricao: "Livre", Preco: 200, Nota: 10},
		{Id: 2, Nome: "Far Cry 5", Descricao: "Livre", Preco: 100, Nota: 9.5},
	}
	log.Println("Conectando Banco de dados...")
	database.ConectaBanco()
	routes.Router()

}
