package controller

import (
	"net/http"

	"github.com/0nF1REy/go-workspace/projects/01_products_api/model"
	"github.com/gin-gonic/gin"
)

// Estrutura do controller
type productController struct {
	// Usecase
}

// Construtor
func NewProductController() productController {
	return productController{}
}

// GET /products
func (p *productController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata Frita",
			Price: 20,
		},
	}

	// Retorna em JSON
	ctx.JSON(http.StatusOK, products)
}
