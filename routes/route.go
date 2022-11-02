package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/loja/controllers"
)

func HandleRoutes() {

	route := gin.Default()

	route.Use()

	//tipoClientes
	route.POST("/api/v1/tipocliente", controllers.InsereTipoCliente)
	route.GET("/api/v1/tipocliente", controllers.BuscaTodosTipoCliente)
	route.GET("/api/v1/tipocliente/:id", controllers.BuscaTipoClientePorID)
	route.DELETE("/api/v1/tipocliente/:id", controllers.ApagaTipoCliente)
	route.PATCH("/api/v1/tipocliente/:id", controllers.AlteraTipoCliente)

	//clientes
	route.POST("api/v1/cliente", controllers.InsereCliente)
	route.GET("api/v1/cliente", controllers.BuscaTodosClientes)
	route.GET("api/v1/cliente/:id", controllers.BuscaClientePorID)
	route.DELETE("api/v1/cliente/:id", controllers.ApagaClientePorID)
	route.PATCH("api/v1/cliente/:id", controllers.AlteraClientePorID)

	//dependentes
	route.POST("api/v1/dependente", controllers.InsereDependente)
	route.GET("api/v1/dependente", controllers.BuscaTodosDependentes)
	route.GET("api/v1/dependente/:id", controllers.BuscaDependentePorID)
	route.GET("api/v1/dependente/cliente/:clienteid", controllers.BuscaDependentePorCliente)
	route.DELETE("api/v1/dependente/:id", controllers.ApagaDependentePorID)
	route.PATCH("api/v1/dependente/:id", controllers.AlteraDependentePorID)

	route.Run() // listen and serve on localhost:8080

}
