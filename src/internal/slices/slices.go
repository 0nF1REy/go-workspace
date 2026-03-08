package slices

import "fmt"

func Execute() {
	fmt.Println("=== Demonstração de Slices em Go ===")

	frutas := []string{"maçã", "banana", "laranja"}
	fmt.Println("Slice inicial:", frutas)

	frutas = append(frutas, "uva")
	fmt.Println("Após append:", frutas)

	numeros := [5]int{10, 20, 30, 40, 50}
	sliceNumeros := numeros[1:4]
	fmt.Println("Slice de array:", sliceNumeros)

	fmt.Println("Len:", len(frutas), "Cap:", cap(frutas))

	fmt.Println("Iterando com for:")
	for i := 0; i < len(frutas); i++ {
		fmt.Println(i, "->", frutas[i])
	}

	fmt.Println("Iterando com for range:")
	for idx, f := range frutas {
		fmt.Printf("%d -> %s\n", idx, f)
	}

	copia := make([]string, len(frutas))
	copy(copia, frutas)
	fmt.Println("Cópia do slice:", copia)

	sub := frutas[1:3]
	fmt.Println("Sub-slice:", sub)

	sub[0] = "morango"
	fmt.Println("Após modificar sub-slice:", frutas)
}
