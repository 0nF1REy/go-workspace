package usecase

import (
	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/model"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/repository"
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

func (pu *ProductUsecase) CreateProduct(p model.Product) (int, error) {
	return pu.Repo.CreateProduct(p)
}

func (pu *ProductUsecase) GetProductById(id_product int) (model.Product, error) {
	return pu.Repo.GetProductById(id_product)
}

func (pu *ProductUsecase) UpdateProduct(id int, p model.Product) error {
	return pu.Repo.UpdateProduct(id, p)
}

func (pu *ProductUsecase) DeleteProduct(id int) error {
	return pu.Repo.DeleteProduct(id)
}
