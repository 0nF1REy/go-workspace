package products_api

import (
	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/controller"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/db"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/middleware"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/repository"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func RunAPI() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// ----------------------
	// MIGRATIONS
	// ----------------------
	if err := db.RunMigrations(dbConnection); err != nil {
		panic(err)
	}

	// ----------------------
	// Camada de repository
	// ----------------------
	productRepo := repository.NewProductRepository(dbConnection)
	authRepo := repository.NewAuthRepository(dbConnection)

	// ----------------------
	// Camada usecase
	// ----------------------
	productUsecase := usecase.NewProductUseCase(productRepo)
	authUsecase, err := usecase.NewAuthUsecaseFromEnv(authRepo)
	if err != nil {
		panic(err)
	}

	// ----------------------
	// Camada de controllers
	// ----------------------
	productController := controller.NewProductController(productUsecase)
	authController := controller.NewAuthController(&authUsecase)

	api := server.Group("/api/v1")

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})
	api.POST("/auth/register", authController.Register)
	api.POST("/auth/login", authController.Login)

	// ----------------------
	// Product API Endpoints
	// ----------------------
	protected := api.Group("")
	protected.Use(middleware.JWTAuthMiddleware(&authUsecase))
	protected.GET("/products", productController.GetProducts)
	protected.POST("/products", productController.CreateProduct)
	protected.GET("/products/:id", productController.GetProductById)
	protected.PUT("/products/:id", productController.UpdateProduct)
	protected.DELETE("/products/:id", productController.DeleteProduct)

	server.Run(":8000")
}
