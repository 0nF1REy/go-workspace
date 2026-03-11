package menu

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func runMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== CLI Interativo [MAIN] ===")
		fmt.Println("1 - Projects")
		fmt.Println("2 - Internal")
		fmt.Println("3 - Examples")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if os.Getenv("DEV_HOT_RELOAD") == "1" {
					fmt.Println("Entrada encerrada (EOF). Modo hot reload detectado: abrindo Projects (01)...")
					Projects(true)
					return
				}
				fmt.Println("Entrada encerrada (EOF). Saindo do menu principal...")
				return
			}
			fmt.Println("Erro ao ler entrada:", err)
			continue
		}

		input = strings.TrimSpace(input)

		switch input {
		case "1":
			Projects(true)
		case "2":
			Internal(true)
		case "3":
			Examples(true)
		case "0":
			fmt.Println("Saindo do menu principal...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}

func StartMenu() {
	runMenu()
}
