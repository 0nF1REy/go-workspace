package main

import "fmt"

// ---------------------------------
// Interface
// Qualquer coisa que tenha Price() e Description()
// pode ser vendida
// ---------------------------------
type Seller interface {
	Price() float64
	Description() string
}

// ---------------------------------
// Album
// ---------------------------------
type Album struct {
	Title     string
	Artist    string
	BasePrice float64
}

func (a Album) Price() float64 {
	return a.BasePrice
}

func (a Album) Description() string {
	return a.Artist + " - " + a.Title
}

// ---------------------------------
// Book
// ---------------------------------
type Book struct {
	Title     string
	Author    string
	BasePrice float64
}

func (b Book) Price() float64 {
	return b.BasePrice
}

func (b Book) Description() string {
	return b.Author + " - " + b.Title
}

// ---------------------------------
// Game
// ---------------------------------
type Game struct {
	Name      string
	Studio    string
	BasePrice float64
}

func (g Game) Price() float64 {
	return g.BasePrice
}

func (g Game) Description() string {
	return g.Studio + " - " + g.Name
}

// ---------------------------------
// Função que vende qualquer item
// ---------------------------------
func SellItem(item Seller) {
	fmt.Println("Item:", item.Description())
	fmt.Println("Preço: R$", item.Price())
	fmt.Println("--------------------------")
}

func main() {

	album := Album{
		Title:     "The Dark Side of the Moon",
		Artist:    "Pink Floyd",
		BasePrice: 59.90,
	}

	book := Book{
		Title:     "Clean Code",
		Author:    "Robert C. Martin",
		BasePrice: 89.90,
	}

	game := Game{
		Name:      "The Witcher 3",
		Studio:    "CD Projekt Red",
		BasePrice: 129.90,
	}

	// Slice de interface
	store := []Seller{album, book, game}

	// Polimorfismo acontecendo aqui
	for _, item := range store {
		SellItem(item)
	}
}
