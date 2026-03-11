package controller

import (
	"net/http"

	"github.com/0nF1REy/go-workspace/projects/01_products_api/usecase"
	"github.com/gin-gonic/gin"
)

// Estrutura do controller
type productController struct {
	productUseCase usecase.ProductUsecase
}

// Construtor
func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

// GET /products
func (p *productController) GetProducts(ctx *gin.Context) {
	// Busca produtos reais do usecase
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}

	// Retorna produtos em JSON
	ctx.JSON(http.StatusOK, products)
}
