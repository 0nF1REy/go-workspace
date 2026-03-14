package controller

import (
	"net/http"
	"strconv"

	"github.com/0nF1REy/go-workspace/projects/01_products_api/model"
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

// POST /products
func (p *productController) CreateProduct(ctx *gin.Context) {
	var input model.Product

	// Bind JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Validação
	if input.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "campo 'name' é obrigatório"})
		return
	}
	if input.Price <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "campo 'price' deve ser maior que zero"})
		return
	}

	// Cria produto
	id, err := p.productUseCase.CreateProduct(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}

	// Retorna produto completo
	product := model.Product{
		ID:    id,
		Name:  input.Name,
		Price: input.Price,
	}

	ctx.JSON(http.StatusCreated, product)
}

// GET /products/:id
func (p *productController) GetProductById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "id do produto deve ser um número",
		})
		return
	}

	product, err := p.productUseCase.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "produto não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
