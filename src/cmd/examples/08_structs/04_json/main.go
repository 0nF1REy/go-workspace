package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   string
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo método:", c.Nome)
}

type ClienteInternacional struct {
	Cliente
	Passaporte string
	Pais       string `json:"pais"`
}

func exibirClientes(clientes []Cliente) {
	fmt.Println("===== LISTA DE CLIENTES =====")

	for _, c := range clientes {
		fmt.Printf("- Nome: %s | Email: %s | CPF: %s\n",
			c.Nome, c.Email, c.CPF)
	}
}

func exibirClientesInternacionais(clientes []ClienteInternacional) {
	fmt.Println("\n===== CLIENTES INTERNACIONAIS =====")

	for _, c := range clientes {
		fmt.Printf("- Nome: %s | Email: %s | CPF: %s | Passaporte: %s | País: %s\n",
			c.Nome,
			c.Email,
			c.CPF,
			c.Passaporte,
			c.Pais,
		)
	}
}

func carregarClientes() ([]Cliente, []ClienteInternacional) {

	cliente := Cliente{
		Nome:  "Alan Ryan",
		Email: "alanryan619@gmail.com",
		CPF:   "4288734234",
	}

	cliente2 := Cliente{
		Nome:  "Bulma",
		Email: "m@gmail.com",
		CPF:   "232323421",
	}

	clienteIntl := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Goku",
			Email: "goku@capsulecorp.com",
			CPF:   "12345678901",
		},
		Passaporte: "ZX123456",
		Pais:       "Japão",
	}

	clientes := []Cliente{cliente, cliente2, clienteIntl.Cliente}
	clientesInternacionais := []ClienteInternacional{clienteIntl}

	return clientes, clientesInternacionais
}

func main() {

	clientes, clientesInternacionais := carregarClientes()

	exibirClientes(clientes)
	exibirClientesInternacionais(clientesInternacionais)

	clientes[0].Exibe()
	clientes[1].Exibe()
	clientes[2].Exibe()

	for _, cliente := range clientes {
		cliente.Exibe()
	}

	clienteJson, err := json.Marshal(clientesInternacionais[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(clienteJson))

	jsonCliente4 := `{"Nome":"Goku","Email":"goku@capsulecorp.com","CPF":"12345678901","Passaporte":"ZX123456","pais":"Japão"}`
	cliente4 := ClienteInternacional{}

	json.Unmarshal([]byte(jsonCliente4), &cliente4)
	fmt.Println(cliente4)
}
