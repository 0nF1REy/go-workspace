package main

import "fmt"

// ---------------------------------
// Interface: algo que pode ser vendido
// ---------------------------------
type Sellable interface {
	Price() float64
	Discount() float64
	Description() string
}

// ---------------------------------
// Struct Album
// ---------------------------------
type Album struct {
	Title        string
	Artist       string
	BasePrice    float64
	DiscountRate float64
}

// Implementação dos métodos da interface
func (a Album) Price() float64 {
	return a.BasePrice
}

func (a Album) Discount() float64 {
	return a.DiscountRate
}

func (a Album) Description() string {
	return a.Artist + " - " + a.Title
}

// ---------------------------------
// Função que vende qualquer item
// ---------------------------------
func SellItem(item Sellable) {

	price := item.Price()
	discount := item.Discount()

	finalPrice := price - (price * discount)

	fmt.Println("Vendendo item:")
	fmt.Println("Descrição:", item.Description())
	fmt.Println("Preço original: R$", price)
	fmt.Println("Desconto:", discount*100, "%")
	fmt.Println("Preço final: R$", finalPrice)
	fmt.Println("---------------------------")
}

func main() {

	album1 := Album{
		Title:        "The Dark Side of the Moon",
		Artist:       "Pink Floyd",
		BasePrice:    59.90,
		DiscountRate: 0.10,
	}

	album2 := Album{
		Title:        "Abbey Road",
		Artist:       "The Beatles",
		BasePrice:    49.90,
		DiscountRate: 0.20,
	}

	SellItem(album1)
	SellItem(album2)
}
