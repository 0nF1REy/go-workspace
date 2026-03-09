package main

import (
	"fmt"
)

// Função que dobra o valor de um inteiro via ponteiro
func doubleValue(num *int) {
	*num = *num * 2
}

func main() {
	// ----------------------------
	// Exemplo básico de ponteiro
	// ----------------------------
	x := 10
	px := &x

	fmt.Println("Endereço de x:", px)
	fmt.Println("Valor de x via ponteiro:", *px)

	// ----------------------------
	// Alterando o valor via ponteiro
	// ----------------------------
	*px = 20
	fmt.Println("Novo valor de x após alteração via ponteiro:", x)

	// ----------------------------
	// Ponteiro passado para função
	// ----------------------------
	y := 5
	fmt.Println("\nAntes da função, y =", y)
	doubleValue(&y)
	fmt.Println("Depois da função, y =", y)

	// ----------------------------
	// Ponteiros com diferentes tipos
	// ----------------------------
	str := "Alan"
	pStr := &str
	fmt.Println("\nEndereço da string:", pStr)
	fmt.Println("Valor da string via ponteiro:", *pStr)
}
