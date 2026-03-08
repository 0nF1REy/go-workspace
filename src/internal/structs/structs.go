package structs

import "fmt"

type Person struct {
    Name   string
    Age    int
    Active bool
}

type Product struct {
    Name   string
    Price  float64
    Stock  int
}

func (p *Person) HaveBirthday() {
    p.Age++
}

func Execute() {
    fmt.Println("=== Demonstração de Structs em Go ===")

    person1 := Person{"Alan Ryan", 20, true}
    person2 := Person{Name: "Carol", Age: 25, Active: false}

    fmt.Println("Pessoa 1:", person1)
    fmt.Println("Pessoa 2:", person2)

    person1.HaveBirthday()
    person2.HaveBirthday()

    fmt.Println("Após fazer aniversário:")
    fmt.Println("Pessoa 1:", person1)
    fmt.Println("Pessoa 2:", person2)

    product := Product{"Caneta", 3.50, 100}
    fmt.Println("Produto:", product)
    product.Stock -= 10
    fmt.Println("Produto após venda:", product)

    people := []Person{person1, person2}
    fmt.Println("Lista de pessoas:", people)

    fmt.Println("\n=== Trabalhando com User e Post ===")

    user := User{ID: 1, Username: "alanryan", Email: "alan@example.com", Active: true}
    fmt.Printf("Usuário: ID=%d, Nome=%s, Email=%s, Ativo=%t\n", user.ID, user.Username, user.Email, user.Active)

    post := Post{ID: 1, Title: "Primeiro Post", Content: "Olá mundo!", Author: user}
    fmt.Printf("Post: ID=%d, Título=%s, Conteúdo=%s, Autor=%s\n", post.ID, post.Title, post.Content, post.Author.Username)
}
