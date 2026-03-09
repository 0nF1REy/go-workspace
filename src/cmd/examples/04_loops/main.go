package main

import "fmt"

func main() {
	// Loop tradicional
	for i := 0; i < 5; i++ {
		fmt.Println("Loop:", i)
	}

	// Range sobre número
	for i := range 3 {
		fmt.Println("Range loop:", i)
	}

	// ----------------------------
	// Iteração sobre string com RANGE
	// ----------------------------
	for i, char := range "Alan" {
		fmt.Println("String range:", i, char)
	}

	// ----------------------------
	// Iteração sobre string com FOR tradicional (bytes)
	// ----------------------------
	str := "Alan"
	for i := 0; i < len(str); i++ {
		c := str[i] // pega o byte no índice i
		fmt.Println("String for (byte):", i, c)
	}

	// ----------------------------
	// Iteração sobre string com FOR tradicional (runes)
	// Para lidar com caracteres Unicode corretamente
	// ----------------------------
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		fmt.Println("String for (rune):", i, runes[i])
	}

	// Loop tipo while
	j := 10
	for j > -1 {
		fmt.Println("While loop:", j)
		j--
	}

	// Loop infinito com break
	for {
		fmt.Println("Loop infinito")
		break
	}

	// 'Continue' pula a iteração atual
	for i := 0; i < 6; i++ {
		if i == 3 {
			fmt.Println("Pulando 3")
			continue
		}
		fmt.Println("Número:", i)
	}

	// Switch simples
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			fmt.Println("Início")
		case 1, 2:
			fmt.Println("Começo ainda:", i)
		case 3:
			fmt.Println("Quase final:", i)
		case 4:
			fmt.Println("Última")
		}
	}
}
