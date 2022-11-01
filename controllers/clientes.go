package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/loja/database"
	"github.com/murilorscampos/loja/models"
)

func BuscaTodosClientes(c *gin.Context) {

	var clientes []models.Cliente

	err := database.DB.Order("nome").Find(&clientes)

	log.Println(err.Error)

	if err != nil {
		panic(err.Error)
	}

	c.JSON(http.StatusOK, clientes)

}
