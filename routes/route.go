package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/loja/controllers"
)

func HandleRoutes() {

	route := gin.Default()

	//tipoClientes
	tiposCliente := route.Group("api/v1/tiposcliente")
	{
		tiposCliente.POST("", controllers.InsereTipoCliente)
		tiposCliente.GET("", controllers.BuscaTodosTipoCliente)
		tiposCliente.GET("/:id", controllers.BuscaTipoClientePorID)
		tiposCliente.DELETE("/:id", controllers.ApagaTipoCliente)
		tiposCliente.PATCH("/:id", controllers.AlteraTipoCliente)
	}

	//clientes
	clientes := route.Group("api/v1/clientes")
	{
		clientes.POST("/", controllers.InsereCliente)
		clientes.GET("/", controllers.BuscaTodosClientes)
		clientes.GET("/:id", controllers.BuscaClientePorID)
		clientes.DELETE("/:id", controllers.ApagaClientePorID)
		clientes.PATCH("/:id", controllers.AlteraClientePorID)
	}

	//dependentes
	dependentes := route.Group("api/v1/dependentes")
	{
		dependentes.POST("", controllers.InsereDependente)
		dependentes.GET("", controllers.BuscaTodosDependentes)
		dependentes.GET("/:id", controllers.BuscaDependentePorID)
		dependentes.GET("/cliente/:clienteid", controllers.BuscaDependentePorCliente)
		dependentes.DELETE("/:id", controllers.ApagaDependentePorID)
		dependentes.PATCH("/:id", controllers.AlteraDependentePorID)
	}

	route.Run() // listen and serve on localhost:8080

}
