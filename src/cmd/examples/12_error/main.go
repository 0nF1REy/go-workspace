package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func fetchWebsiteHeaders() {

	res, err := http.Get("https://alan-ryan.vercel.app")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res.Header)
}

func sumWithLimit(x int, y int) (int, error) {

	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}

func handleSum() {
	res, err := sumWithLimit(10, 10)

	if err != nil {
		log.Println("Erro ao somar:", err)
		return
	}

	fmt.Println("Resultado:", res)
}

func ignoreError() {
	res, _ := sumWithLimit(3, 4)
	fmt.Println("Resultado sem tratar erro:", res)
}

func main() {
	fetchWebsiteHeaders()
	handleSum()
	ignoreError()
}
