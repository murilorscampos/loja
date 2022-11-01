package database

import (
	"log"

	"github.com/murilorscampos/loja/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBancoDados() {

	stringDeConexao := "host=localhost user=postgres password=root dbname=loja port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar no banco de dados.")
	} else {
		log.Println("Banco de dados conectado...")
	}

	DB.AutoMigrate(
		&models.TipoCliente{},
		&models.Cliente{},
		&models.Dependente{},
	)

}
