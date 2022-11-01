package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/loja/database"
	"github.com/murilorscampos/loja/models"
)

func InsereTipoCliente(c *gin.Context) {

	var tipoCliente models.TipoCliente

	if err := c.ShouldBindJSON(&tipoCliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error()})

		return
	}

	if err := models.ValidaDadosDeAluno(&tipoCliente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error()})

		return
	}

	database.DB.Create(&tipoCliente)
	c.JSON(http.StatusOK, tipoCliente)

}

func BuscaTodosTiposClientes(c *gin.Context) {

	var tiposClientes []models.TipoCliente

	database.DB.Order("descricao").Find(&tiposClientes)

	c.JSON(http.StatusOK, tiposClientes)

}

func BuscaTipoClientePorID(c *gin.Context) {

	var tipoCliente models.TipoCliente

	id := c.Params.ByName("id")

	linhas := database.DB.First(&tipoCliente, id).RowsAffected

	if linhas == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"data": "Tipo de cliente não encontrado, ID " + string(id),
		})

		return
	}

	c.JSON(http.StatusOK, tipoCliente)

}

func ApagaTipoCliente(c *gin.Context) {

	var tipoCliente models.TipoCliente

	id := c.Params.ByName("id")

	linhas := database.DB.Delete(&tipoCliente, id).RowsAffected

	if linhas == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"data": "Tipo de cliente não encontrado, ID " + string(id),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Tipo de cliente excluído com sucesso, ID " + id,
	})

}
