package main

import "fmt"

// -------------------------------
// Definição da struct Person
// -------------------------------
type Person struct {
	Name   string
	Age    int
	City   string
	Active bool
}

// Função que cria uma nova Person (retornando valor)
func newPersonValue(name string, age int, city string, active bool) Person {
	return Person{
		Name:   name,
		Age:    age,
		City:   city,
		Active: active,
	}
}

// Função que tenta modificar a struct (por valor)
func modifyValue(p Person) {
	p.Name = "Alan" // NÃO altera o original
}

// Exibe os dados da pessoa
func displayPerson(p Person) {
	fmt.Println("===================================")
	fmt.Println("          DETALHES DA PESSOA       ")
	fmt.Println("===================================")
	fmt.Println("Nome  :", p.Name)
	fmt.Println("Idade :", p.Age)
	fmt.Println("Cidade:", p.City)
	fmt.Println("Ativo :", p.Active)
	fmt.Println("===================================")
}

func main() {

	alice := newPersonValue("Alice", 25, "São Paulo", true)
	bob := newPersonValue("Bob", 30, "Rio de Janeiro", false)

	displayPerson(bob)

	fmt.Println("Antes da modificação:", alice.Name)

	modifyValue(alice)

	fmt.Println("Depois da modificação:", alice.Name)
}
