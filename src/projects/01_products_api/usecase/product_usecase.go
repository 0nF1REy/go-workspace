package usecase

import (
	"github.com/0nF1REy/go-workspace/projects/01_products_api/model"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/repository"
)

type ProductUsecase struct {
	Repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{Repo: repo}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.Repo.GetProducts()
}
