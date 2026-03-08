package types

import "fmt"

func Execute() {
    var idade int = 20
    var ano int64 = 2026

    var preco float64 = 99.99
    var taxa float32 = 0.15

    var ativo bool = true

    var nome string = "Alan Ryan"

    var numeros [3]int = [3]int{1, 2, 3}

    frutas := []string{"maçã", "banana", "laranja"}

    notas := map[string]float64{
        "Matemática": 9.5,
        "Física":     8.0,
        "Química":    7.5,
    }

    fmt.Println("Inteiros:", idade, ano)
    fmt.Println("Float:", preco, taxa)
    fmt.Println("Boolean:", ativo)
    fmt.Println("String:", nome)
    fmt.Println("Array:", numeros)
    fmt.Println("Slice:", frutas)
    fmt.Println("Map:", notas)
}
