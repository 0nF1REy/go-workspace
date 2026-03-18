// Auto-pagination: segue automaticamente os links de "next page"

package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector()

	// Quando encontrar o botão "next", visita automaticamente a próxima página
	c.OnHTML("li.next > a", func(e *colly.HTMLElement) {
		nextPageURL := e.Request.AbsoluteURL(e.Attr("href"))
		if nextPageURL == "https://sandbox.oxylabs.io/products?page=6" {
			return
		}
		fmt.Println("Encontrado 'next page'", nextPageURL)
		e.Request.Visit(nextPageURL)
	})

	url := "https://sandbox.oxylabs.io/products"
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
