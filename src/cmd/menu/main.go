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

func StartMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== CLI Interativo Go Workspace ===")
		fmt.Println("1 - Tipos básicos")
		fmt.Println("2 - Estruturas de controle")
		fmt.Println("3 - Funções")
		fmt.Println("4 - Slices")
		fmt.Println("5 - Structs")
		fmt.Println("6 - API de Produtos")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		input, _ := reader.ReadString('\n')
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
		case "6":
			fmt.Println("Iniciando API de Produtos...")
			products_api.StartServer()
			return
		case "0":
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
