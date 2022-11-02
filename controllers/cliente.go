package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/loja/database"
	"github.com/murilorscampos/loja/models"
)

func InsereCliente(c *gin.Context) {

	var cliente models.Cliente

	if err := c.ShouldBindJSON(&cliente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		return

	}

	if err := models.ValidaDadosCliente(&cliente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		return

	}

	if err := database.DB.Create(&cliente); err.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error,
		})

		return

	}

	c.JSON(http.StatusOK, cliente)

}

func BuscaTodosClientes(c *gin.Context) {

	var clientes []models.Cliente

	if err := database.DB.Order("nome").Find(&clientes); err.Error != nil {

		c.JSON(http.StatusBadRequest, err.Error)

		return

	}

	c.JSON(http.StatusOK, clientes)

}

func BuscaClientePorID(c *gin.Context) {

	var cliente models.Cliente

	id := c.Params.ByName("id")

	if result := database.DB.First(&cliente, id); result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Cliente não encontrado, ID " + id,
		})

		return

	}

	c.JSON(http.StatusOK, cliente)

}

func ApagaClientePorID(c *gin.Context) {

	var cliente models.Cliente

	id := c.Params.ByName("id")

	if result := database.DB.Delete(&cliente, id); result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Cliente não encontrado, ID " + id,
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Cliente excluído com sucesso, ID: " + id,
	})

}

func AlteraClientePorID(c *gin.Context) {

	var cliente models.Cliente

	id := c.Params.ByName("id")

	if result := database.DB.First(&cliente, id); result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Cliente não encontrado, ID " + id,
		})

		return

	}

	if err := c.ShouldBindJSON(&cliente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		return

	}

	if err := models.ValidaDadosCliente(&cliente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		return

	}

	if err := database.DB.UpdateColumns(&cliente); err.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error,
		})

		return

	}

	c.JSON(http.StatusOK, cliente)

}
