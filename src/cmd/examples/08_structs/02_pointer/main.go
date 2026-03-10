package main

import "fmt"

// Definição de uma struct chamada 'Person'
type Person struct {
	Name   string
	Age    int
	City   string
	Active bool
}

// Criação de uma instância de Person (ponteiro)
func newPersonPointer(name string, age int, city string, active bool) *Person {
	return &Person{
		Name:   name,
		Age:    age,
		City:   city,
		Active: active,
	}
}

// Método da struct para exibir
func (p *Person) display(short bool) {
	if short {
		fmt.Println("Desenvolvedor:", p.Name)
	} else {
		fmt.Printf("Nome: %s, Idade: %d, Cidade: %s, Ativo: %t\n", p.Name, p.Age, p.City, p.Active)
	}
}

// Modifica a struct e chama display
func modifyPointer(p *Person, short bool) {
	p.Name = "Alan"
	p.display(true)
}

func main() {
	alice := newPersonPointer("Alice", 25, "São Paulo", true)
	fmt.Println("Antes da modificação:", alice.Name)

	modifyPointer(alice, true) // aqui passamos o booleano
	fmt.Println("Depois da modificação (ponteiro):", alice.Name)
}
