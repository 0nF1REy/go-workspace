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
// Função que identifica o tipo real
// ---------------------------------
func CheckItemType(item Seller) {

	fmt.Println("Item:", item.Description())

	// Type assertion
	if album, ok := item.(Album); ok {
		fmt.Println("Tipo detectado: Album")
		fmt.Println("Artista:", album.Artist)
	}

	if book, ok := item.(Book); ok {
		fmt.Println("Tipo detectado: Book")
		fmt.Println("Autor:", book.Author)
	}

	fmt.Println("Preço:", item.Price())
	fmt.Println("-------------------------")
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

	items := []Seller{album, book}

	for _, item := range items {
		CheckItemType(item)
	}
}
