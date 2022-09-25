package main

import (
	"App/database"
	"App/routes"
	"log"
)

func main() {
	log.Println("inicializando Servidor...")
	database.ConectaBanco()
	routes.Router()

}
