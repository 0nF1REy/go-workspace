package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

type Product struct {
	Name  string
	Price string
	URL   string
}

func main() {

	c := colly.NewCollector()
	var products []Product

	c.OnHTML("div.product-card", func(e *colly.HTMLElement) {
		product := Product{
			Name:  e.ChildText("h4.title"),
			Price: e.ChildText("div.price-wrapper"),
			URL:   e.Request.AbsoluteURL(e.ChildAttr("a.card-header", "href")),
		}
		products = append(products, product)
	})

	c.OnScraped(func(r *colly.Response) {
		for _, p := range products {
			fmt.Printf("Nome: %s, Preço: %s, URL: %s\n", p.Name, p.Price, p.URL)
		}
	})

	err := c.Visit("https://sandbox.oxylabs.io/products")
	if err != nil {
		log.Fatal(err)
	}

	// ==============================
	// Chama a função do 'export.go'
	// ==============================
	if err := ExportCSV(products, "products.csv"); err != nil {
		log.Fatal("Erro ao exportar CSV:", err)
	}
}
