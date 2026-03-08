package control

import "fmt"

func Execute() {
    fmt.Println("=== Estruturas de Controle em Go ===")

    numero := 7
    if numero%2 == 0 {
        fmt.Println(numero, "é par")
    } else {
        fmt.Println(numero, "é ímpar")
    }

    fmt.Println("\nFor clássico:")
    for i := 0; i < 5; i++ {
        fmt.Println("i =", i)
    }

    fmt.Println("\nFor como while:")
    contador := 0
    for contador < 3 {
        fmt.Println("contador =", contador)
        contador++
    }

    fmt.Println("\nFor range com slice:")
    frutas := []string{"maçã", "banana", "laranja"}
    for idx, fruta := range frutas {
        fmt.Printf("%d -> %s\n", idx, fruta)
    }

    fmt.Println("\nFor range com map:")
    pessoas := map[string]int{"Alice": 25, "Bob": 30, "Carol": 22}
    for nome, idade := range pessoas {
        fmt.Printf("%s tem %d anos\n", nome, idade)
    }

    fmt.Println("\nSwitch example:")
    letra := "b"
    switch letra {
    case "a":
        fmt.Println("É A")
    case "b":
        fmt.Println("É B")
    default:
        fmt.Println("Não é A nem B")
    }
}
