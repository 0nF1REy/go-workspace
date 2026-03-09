package main

import "fmt"

func main() {
	fmt.Println("- - - - - - - - - - -")
	// ----------------------------
	// Array de 5 inteiros
	// ----------------------------
	var numeros [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array completo:", numeros)

	// Acesso a elementos individuais
	fmt.Println("Primeiro elemento:", numeros[0])
	fmt.Println("Último elemento:", numeros[len(numeros)-1])
	fmt.Println("- - - - - - - - - - -")

	// Iteração usando for tradicional
	for i := 0; i < len(numeros); i++ {
		fmt.Println("Indice", i, "valor:", numeros[i])
	}

	fmt.Println("- - - - - - - - - - -")

	// Iteração usando for range
	for i, valor := range numeros {
		fmt.Println("Range -> indice:", i, "valor:", valor)
	}

	fmt.Println("= = = = = = = = = = =")

	// ----------------------------
	// Slice e append
	// ----------------------------
	var arr1 []int
	arr1 = append(arr1, 5)         
	arr1 = append(arr1, 100, 300, 200, -666)

	// Iteração sobre slice
	for i, valor := range arr1 {
		fmt.Println("Slice -> indice:", i, "valor:", valor)
	}

	fmt.Println("Slice após append:", arr1)
}
