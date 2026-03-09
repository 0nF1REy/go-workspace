package main

import "fmt"

var albumName string
var albumPrice int = 50
var artistName = "Danny Joe Brown Band"

func main() {

	const availableStock = 50

	genre := "Southern Rock"

	fmt.Println("Nome do álbum:", albumName+" |", "Preço:", albumPrice, "|", "Artista/Banda:", artistName)
	fmt.Println("Gênero:", genre)
	fmt.Println("Estoque disponível:", availableStock)
}
