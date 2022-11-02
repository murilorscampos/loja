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
			"data": err.Error(),
		})

		return

	}

	if err := models.ValidaDadosTipoCliente(&tipoCliente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		return

	}

	if result := database.DB.Create(&tipoCliente); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	}

	c.JSON(http.StatusOK, tipoCliente)

}

func BuscaTodosTipoCliente(c *gin.Context) {

	var tiposClientes []models.TipoCliente

	if result := database.DB.Order("descricao").Find(&tiposClientes); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	}

	c.JSON(http.StatusOK, tiposClientes)

}

func BuscaTipoClientePorID(c *gin.Context) {

	var tipoCliente models.TipoCliente

	id := c.Params.ByName("id")

	if result := database.DB.First(&tipoCliente, id); result.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": "Tipo cliente não encontrado, ID " + id,
		})

		return

	}

	c.JSON(http.StatusOK, tipoCliente)

}

func ApagaTipoCliente(c *gin.Context) {

	var tipoCliente models.TipoCliente

	id := c.Params.ByName("id")

	result := database.DB.Delete(&tipoCliente, id)

	if result.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": "Tipo cliente não encontrado, ID " + id,
		})

		return

	} else {

		if result.Error != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"data": result.Error,
			})

			return

		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Tipo de cliente excluído com sucesso, ID " + id,
	})

}

func AlteraTipoCliente(c *gin.Context) {

	var tipoCliente models.TipoCliente

	id := c.Params.ByName("id")

	result := database.DB.First(&tipoCliente, id)

	if result.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": "Tipo cliente não encontrado, ID " + id,
		})

		return

	} else {

		if result.Error != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"data": result.Error,
			})

			return

		}
	}

	if err := c.ShouldBindJSON(&tipoCliente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		return

	}

	if err := models.ValidaDadosTipoCliente(&tipoCliente); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error()})

		return

	}

	if err := database.DB.UpdateColumns(&tipoCliente); err.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	}

	c.JSON(http.StatusOK, tipoCliente)

}
