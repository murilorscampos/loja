package main

import (
	"log"

	"github.com/murilorscampos/loja/database"
	"github.com/murilorscampos/loja/routes"
)

func main() {

	log.Println("Conectando ao banco de dados...")
	database.ConectaBancoDados()

	routes.HandleRoutes()

}
