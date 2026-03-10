package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/0nF1REy/go-workspace/internal/control"
	"github.com/0nF1REy/go-workspace/internal/functions"
	"github.com/0nF1REy/go-workspace/internal/slices"
	"github.com/0nF1REy/go-workspace/internal/structs"
	"github.com/0nF1REy/go-workspace/internal/types"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/cmd"
)

func Projects(enable bool) {
	if !enable {
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== CLI Interativo [PROJECTS] ===")
		fmt.Println("1 - API de Produtos")
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
			fmt.Println("Iniciando API de Produtos...")
			products_api.StartServer()
			return
		case "0":
			fmt.Println("Saindo do menu Projects...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}

func Internal(enable bool) {
	if !enable {
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== CLI Interativo [INTERNAL] ===")
		fmt.Println("1 - Tipos básicos")
		fmt.Println("2 - Estruturas de controle")
		fmt.Println("3 - Funções")
		fmt.Println("4 - Slices")
		fmt.Println("5 - Structs")
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
			types.Execute()
		case "2":
			control.Execute()
		case "3":
			functions.Execute()
		case "4":
			slices.Execute()
		case "5":
			structs.Execute()
		case "0":
			fmt.Println("Saindo do menu Internal...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}

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

// Função privada para controlar quais menus serão executados
func runMenu(projectsEnabled, internalEnabled, examplesEnabled bool) {
	Projects(projectsEnabled)
	Internal(internalEnabled)
	Examples(examplesEnabled)
}

// Função pública para iniciar o menu
func StartMenu() {
	runMenu(true, false, false)
}
