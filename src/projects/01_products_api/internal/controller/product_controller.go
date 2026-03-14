package controller

import (
	"net/http"
	"strconv"

	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/model"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/usecase"
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
	var input model.CreateProductRequest

	// Bind JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	productInput := model.Product{
		Name:  input.Name,
		Price: input.Price,
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
	id, err := p.productUseCase.CreateProduct(productInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}

	// Retorna produto completo
	product := model.Product{
		ID:    id,
		Name:  productInput.Name,
		Price: productInput.Price,
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

// PUT /products/:id
func (p *productController) UpdateProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "id do produto deve ser um número",
		})
		return
	}

	var input model.Product

	var request model.UpdateProductRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "corpo JSON inválido",
		})
		return
	}

	input = model.Product{
		Name:  request.Name,
		Price: request.Price,
	}

	// Validação
	if input.Name == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "campo 'name' é obrigatório",
		})
		return
	}

	if input.Price <= 0 {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "campo 'price' deve ser maior que zero",
		})
		return
	}

	err = p.productUseCase.UpdateProduct(id, input)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "produto não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "produto atualizado com sucesso",
	})
}

// DELETE /products/:id
func (p *productController) DeleteProduct(ctx *gin.Context) {

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "id do produto deve ser um número",
		})
		return
	}

	err = p.productUseCase.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "produto não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "produto removido com sucesso",
	})
}
