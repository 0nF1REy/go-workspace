package main

import "fmt"

func main() {
	// ----------------------------
	// Declaração e inicialização de um map
	// ----------------------------
	productPrices := map[string]float64{
		"Notebook": 3500.00,
		"Mouse":    50.00,
		"Teclado":  120.00,
	}

	fmt.Println("Map completo:", productPrices)

	// ----------------------------
	// Acessando valores pelo key
	// ----------------------------
	fmt.Println("Preço do Notebook:", productPrices["Notebook"])

	// ----------------------------
	// Adicionando um novo item
	// ----------------------------
	productPrices["Monitor"] = 800.00
	fmt.Println("Map atualizado:", productPrices)

	// ----------------------------
	// Iterando sobre o map
	// ----------------------------
	for produto, preco := range productPrices {
		fmt.Println("Produto:", produto, "| Preço:", preco)
	}

	// ----------------------------
	// Verificando se uma chave existe
	// ----------------------------
	preco, ok := productPrices["Mouse"]
	if ok {
		fmt.Println("Preço do Mouse:", preco)
	} else {
		fmt.Println("Produto não encontrado")
	}

	// ----------------------------
	// Deletando um item
	// ----------------------------
	delete(productPrices, "Teclado")
	fmt.Println("Map após remover Teclado:", productPrices)
}
