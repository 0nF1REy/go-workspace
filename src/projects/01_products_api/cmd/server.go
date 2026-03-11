package products_api

import (
	"github.com/0nF1REy/go-workspace/projects/01_products_api/controller"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/db"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/repository"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/usecase"
	"github.com/gin-gonic/gin"
)

func RunAPI() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	productRepo := repository.NewProductRepository(dbConnection)

	// Camada usecase
	productUsecase := usecase.NewProductUseCase(productRepo)

	// Camada de controllers
	productController := controller.NewProductController(productUsecase)

	api := server.Group("/api/v1")

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	api.GET("/products", productController.GetProducts)
	api.POST("/products", productController.CreateProduct)

	server.Run(":8000")
}
