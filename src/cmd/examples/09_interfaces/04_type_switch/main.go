package main

import "fmt"

// ---------------------------------
// Interface
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
// Função usando type switch
// ---------------------------------
func IdentifyItem(item Seller) {

	fmt.Println("Item:", item.Description())

	switch v := item.(type) {

	case Album:
		fmt.Println("Tipo detectado: Album")
		fmt.Println("Artista:", v.Artist)

	case Book:
		fmt.Println("Tipo detectado: Book")
		fmt.Println("Autor:", v.Author)

	case Game:
		fmt.Println("Tipo detectado: Game")
		fmt.Println("Studio:", v.Studio)

	default:
		fmt.Println("Tipo desconhecido")
	}

	fmt.Println("Preço:", item.Price())
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

	items := []Seller{album, book, game}

	for _, item := range items {
		IdentifyItem(item)
	}
}
