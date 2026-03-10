package main

import "fmt"

// ---------------------------------
// Função que recebe qualquer tipo
// ---------------------------------
func PrintValue(value interface{}) {

	fmt.Println("Valor recebido:", value)

	// Descobrindo o tipo com type switch
	switch v := value.(type) {

	case string:
		fmt.Println("Tipo detectado: string")
		fmt.Println("Conteúdo:", v)

	case int:
		fmt.Println("Tipo detectado: int")
		fmt.Println("Conteúdo:", v)

	case bool:
		fmt.Println("Tipo detectado: bool")
		fmt.Println("Conteúdo:", v)

	case float64:
		fmt.Println("Tipo detectado: float64")
		fmt.Println("Conteúdo:", v)

	default:
		fmt.Println("Tipo desconhecido")
	}

	fmt.Println("---------------------------")
}

func main() {

	values := []interface{}{
		"Hello Go",
		42,
		true,
		3.14,
	}

	for _, v := range values {
		PrintValue(v)
	}
}
