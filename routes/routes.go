package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/loja/controllers"
)

func HandleRoutes() {

	route := gin.Default()
	//tipoClientes
	route.POST("/tipoclientes", controllers.InsereTipoCliente)
	route.GET("/tipoclientes", controllers.BuscaTodosTiposClientes)
	route.GET("/tipoclientes/:id", controllers.BuscaTipoClientePorID)
	route.DELETE("/tipoclientes/:id", controllers.ApagaTipoCliente)

	//clientes
	route.GET("/clientes", controllers.BuscaTodosClientes)
	route.Run() // listen and serve on localhost:8080

}
