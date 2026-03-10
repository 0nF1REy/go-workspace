package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Examples(enable bool) {
	if !enable {
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== CLI Interativo [EXAMPLES] ===")
		fmt.Println("1 - Mostrar Hello World")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			continue
		}
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("Hello World!")
		case "0":
			fmt.Println("Saindo do menu Examples...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
