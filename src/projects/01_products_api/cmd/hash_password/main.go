package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("uso: go run ./projects/01_products_api/cmd/hash_password <senha>")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(os.Args[1]), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("erro ao gerar hash: %v", err)
	}

	fmt.Println(string(hash))
}
