package repository

import (
	"database/sql"
	"log"

	"github.com/0nF1REy/go-workspace/projects/01_products_api/model"
)

type ProductRepository struct {
	Connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		Connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"

	rows, err := pr.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			log.Println("Erro ao escanear produto:", err)
			continue
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *ProductRepository) CreateProduct(p model.Product) (int, error) {
	query := `
		INSERT INTO product (product_name, price)
		VALUES ($1, $2)
		RETURNING id
	`

	var id int

	err := pr.Connection.QueryRow(query, p.Name, p.Price).Scan(&id)
	if err != nil {
		log.Println("Erro ao inserir produto:", err)
		return 0, err
	}

	return id, nil
}
