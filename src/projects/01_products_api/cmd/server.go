package products_api

import (
	"github.com/0nF1REy/go-workspace/projects/01_products_api/controller"
	"github.com/gin-gonic/gin"
)

func Server() {
	ProductController := controller.NewProductController()

	server := gin.Default()

	// Rota de teste
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Rota de produtos
	server.GET("/products", ProductController.GetProducts)

	// Inicia o servidor na porta 8000
	server.Run(":8000")
}
