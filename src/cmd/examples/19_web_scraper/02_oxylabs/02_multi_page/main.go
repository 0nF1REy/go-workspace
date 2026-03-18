// Multipage = Várias requisições

package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

type Product struct {
	Name  string
	Price string
	URL   string
}

func main() {

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitando", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Resposta de", r.Request.URL.String())
	})

	urls := []string{
		"https://sandbox.oxylabs.io/products/1",
		"https://sandbox.oxylabs.io/products/2",
		"https://sandbox.oxylabs.io/products/3",
		"https://sandbox.oxylabs.io/products/4",
	}

	for _, url := range urls {
		c.Visit(url)
	}
	c.Wait()
}
