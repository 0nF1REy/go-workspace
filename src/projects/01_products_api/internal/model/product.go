package model

type Product struct {
	ID    int     `db:"id_product" json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
