// Auto-pagination: segue automaticamente os links de "next page"

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector()

	c.SetRequestTimeout(10 * time.Second)

	// Quando encontrar o botão "next", visita automaticamente a próxima página
	c.OnHTML("li.next > a", func(e *colly.HTMLElement) {
		nextPageURL := e.Request.AbsoluteURL(e.Attr("href"))
		if nextPageURL == "https://sandbox.oxylabs.io/products?page=6" {
			return
		}
		fmt.Println("Encontrado 'next page'", nextPageURL)
		e.Request.Visit(nextPageURL)
	})

	// Tratamento de erro
	c.OnError(func(r *colly.Response, err error) {
		switch r.StatusCode {

		case http.StatusNotFound:
			fmt.Println("Erro 404: Página não encontrada:", r.Request.URL)

		case http.StatusTooManyRequests:
			fmt.Println("Erro 429: Muitas requisições (rate limit atingido). Tente novamente mais tarde:", r.Request.URL)

		case http.StatusInternalServerError:
			fmt.Println("Erro 500: Erro interno do servidor:", r.Request.URL)

		default:
			fmt.Printf("Erro %d ao acessar %s: %v\n", r.StatusCode, r.Request.URL, err)
		}
	})

	// url := "https://sandbox.oxylabs.io/products/non-existent-page"
	url := "https://sandbox.oxylabs.io/products"

	err := c.Visit(url)

	if err != nil {
		log.Fatal(err)
	}
}
