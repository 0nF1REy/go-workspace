package main

import "fmt"

func getEvenOdd(num int) (string, int) {
	if num%2 == 0 {
		return "Par", num
	} else {
		return "Ímpar", num
	}
}

func interactiveEvenOdd() {
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	resultado, valor := getEvenOdd(numero)
	fmt.Printf("O número %d é %s\n", valor, resultado)
}

func main() {
	// Chamada interativa
	interactiveEvenOdd()

	// Chamada direta
	x, y := getEvenOdd(11)
	fmt.Printf("Chamando diretamente: Número %d é %s\n", y, x)
}
