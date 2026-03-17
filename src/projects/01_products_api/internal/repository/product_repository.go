package repository

import (
	"database/sql"
	"log"

	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/model"
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
	query := "SELECT id, name, price FROM products.product"

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
	INSERT INTO products.product (name, price)
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

func (pr *ProductRepository) GetProductById(id_product int) (model.Product, error) {
	query := "SELECT id, name, price FROM products.product WHERE id = $1"

	var p model.Product

	err := pr.Connection.QueryRow(query, id_product).Scan(
		&p.ID,
		&p.Name,
		&p.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Produto não encontrado")
			return model.Product{}, err
		}

		log.Println("Erro ao buscar produto:", err)
		return model.Product{}, err
	}

	return p, nil
}

func (pr *ProductRepository) UpdateProduct(id int, p model.Product) error {

	query := `
	UPDATE products.product
	SET name = $1, price = $2
	WHERE id = $3
	`

	result, err := pr.Connection.Exec(query, p.Name, p.Price, id)
	if err != nil {
		log.Println("Erro ao atualizar produto:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {

	query := "DELETE FROM products.product WHERE id = $1"

	result, err := pr.Connection.Exec(query, id)
	if err != nil {
		log.Println("Erro ao deletar produto:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
