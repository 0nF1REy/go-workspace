package main

import "fmt"

// -------------------------------
// Definição de uma struct chamada 'Person'
// -------------------------------
type Person struct {
	Name   string
	Age    int
	City   string
	Active bool
}

// -------------------------------
// Função que cria uma nova Person (ponteiro)
// -------------------------------
func newPersonPointer(name string, age int, city string, active bool) *Person {
	return &Person{
		Name:   name,
		Age:    age,
		City:   city,
		Active: active,
	}
}

// -------------------------------
// Método da struct para exibir
// -------------------------------
func (p *Person) display(short bool) {
	if short {
		fmt.Println("Nome:", p.Name)
	} else {
		fmt.Printf("Nome: %s, Idade: %d, Cidade: %s, Ativo: %t\n",
			p.Name, p.Age, p.City, p.Active)
	}
}

// -------------------------------
// Função para atualizar idade de cada pessoa
// -------------------------------
func incrementAges(people []*Person) {
	for _, person := range people {
		person.Age += 1 // Modifica direto o objeto original (porque é ponteiro)
	}
}

// -------------------------------
// Função para exibir todas as pessoas
// -------------------------------
func displayPeople(people []*Person, short bool) {
	for i, person := range people {
		fmt.Printf("Pessoa %d: ", i+1)
		person.display(short)
	}
}

func main() {
	// Criando structs como ponteiros
	alice := newPersonPointer("Alice", 25, "São Paulo", true)
	bob := newPersonPointer("Bob", 30, "Rio de Janeiro", false)
	charlie := newPersonPointer("Charlie", 22, "Belo Horizonte", true)

	// Slice de ponteiros
	people := []*Person{alice, bob, charlie}

	fmt.Println("Antes da atualização de idades:")
	displayPeople(people, false) // false = mostra todos os detalhes

	// Atualizando idades diretamente via ponteiros
	incrementAges(people)

	fmt.Println("\nDepois da atualização de idades:")
	displayPeople(people, false)

	fmt.Println("\nExibindo apenas nomes (short = true):")
	displayPeople(people, true)
}
