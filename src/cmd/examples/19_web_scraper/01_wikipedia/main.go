package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector()

	var links []string
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			links = append(links, link)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Encontrado %d links: \n\n", len(links))
		for i, link := range links {
			fmt.Printf("%3d. %s\n", i+1, link)
		}
	})

	err := c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
	if err != nil {
		log.Fatal(err)
	}
}
