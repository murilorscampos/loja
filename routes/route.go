package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/loja/controllers"
)

func HandleRoutes() {

	route := gin.Default()

	route.Use()

	//tipoClientes
	route.POST("/api/v1/tiposcliente", controllers.InsereTipoCliente)
	route.GET("/api/v1/tiposcliente", controllers.BuscaTodosTipoCliente)
	route.GET("/api/v1/tiposcliente/:id", controllers.BuscaTipoClientePorID)
	route.DELETE("/api/v1/tiposcliente/:id", controllers.ApagaTipoCliente)
	route.PATCH("/api/v1/tiposcliente/:id", controllers.AlteraTipoCliente)

	//clientes
	route.POST("api/v1/clientes", controllers.InsereCliente)
	route.GET("api/v1/clientes", controllers.BuscaTodosClientes)
	route.GET("api/v1/clientes/:id", controllers.BuscaClientePorID)
	route.DELETE("api/v1/clientes/:id", controllers.ApagaClientePorID)
	route.PATCH("api/v1/clientes/:id", controllers.AlteraClientePorID)

	//dependentes
	route.POST("api/v1/dependentes", controllers.InsereDependente)
	route.GET("api/v1/dependentes", controllers.BuscaTodosDependentes)
	route.GET("api/v1/dependentes/:id", controllers.BuscaDependentePorID)
	route.GET("api/v1/dependentes/cliente/:clienteid", controllers.BuscaDependentePorCliente)
	route.DELETE("api/v1/dependentes/:id", controllers.ApagaDependentePorID)
	route.PATCH("api/v1/dependentes/:id", controllers.AlteraDependentePorID)

	route.Run() // listen and serve on localhost:8080

}
