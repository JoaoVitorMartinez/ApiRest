package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	conexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conexao))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados!")
	}
	log.Println("Banco de dados conectado.")
}
