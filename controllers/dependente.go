package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/loja/database"
	"github.com/murilorscampos/loja/models"
)

func InsereDependente(c *gin.Context) {

	var dependente models.Dependente

	if err := c.ShouldBindJSON(&dependente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		return

	}

	if err := models.ValidaDadosDependente(&dependente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		return

	}

	if err := database.DB.Create(&dependente); err.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error,
		})

		return

	}

	c.JSON(http.StatusOK, dependente)

}

func BuscaTodosDependentes(c *gin.Context) {

	var dependentes []models.Dependente

	if err := database.DB.Order("nome").Find(&dependentes); err.Error != nil {

		c.JSON(http.StatusBadRequest, err.Error)

		return

	}

	c.JSON(http.StatusOK, dependentes)

}

func BuscaDependentePorID(c *gin.Context) {

	var dependente models.Dependente

	id := c.Params.ByName("id")

	if result := database.DB.Order("nome").Find(&dependente, id); result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Dependente não encontrado, ID " + id,
		})

		return

	}

	c.JSON(http.StatusOK, dependente)

}

func BuscaDependentePorCliente(c *gin.Context) {

	var dependentes []models.Dependente

	clienteId, _ := strconv.Atoi(c.Params.ByName("clienteid"))

	if result := database.DB.Where(&models.Dependente{ClienteID: clienteId}).Find(&dependentes); result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Dependente não encontrado para esse cliente, ClienteID " + strconv.Itoa(clienteId),
		})

		return

	}

	c.JSON(http.StatusOK, dependentes)

}

func ApagaDependentePorID(c *gin.Context) {

	var dependente models.Dependente

	id := c.Params.ByName("id")

	if result := database.DB.Delete(&dependente, id); result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Dependente não encontrado, ID " + id,
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Dependente excluído com sucesso, ID: " + id,
	})

}

func AlteraDependentePorID(c *gin.Context) {

	var dependente models.Dependente

	id := c.Params.ByName("id")

	if result := database.DB.First(&dependente, id); result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Dependente não encontrado, ID " + id,
		})

		return

	}

	if err := c.ShouldBindJSON(&dependente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		return

	}

	if err := models.ValidaDadosDependente(&dependente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		return

	}

	if err := database.DB.UpdateColumns(&dependente); err.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error,
		})

		return

	}

	c.JSON(http.StatusOK, dependente)

}
