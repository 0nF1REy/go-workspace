package menu

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	products_api "github.com/0nF1REy/go-workspace/projects/01_products_api/cmd"
	tasks_api "github.com/0nF1REy/go-workspace/projects/02_tasks_api/cmd"
)

func Projects(enable bool) {
	if !enable {
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== CLI Interativo [PROJECTS] ===")
		fmt.Println("1 - API de Produtos")
		fmt.Println("2 - API de Tarefas")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if os.Getenv("DEV_HOT_RELOAD") == "1" {
					fmt.Println("Entrada encerrada (EOF). Iniciando API de Produtos no modo hot reload...")
					products_api.StartAPI()
					return
				}
				fmt.Println("Entrada encerrada (EOF). Saindo do menu Projects...")
				return
			}
			fmt.Println("Erro ao ler entrada:", err)
			continue
		}
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("Iniciando API de Produtos...")
			products_api.StartAPI()
			return
		case "2":
			fmt.Println("Iniciando API de Tarefas...")
			tasks_api.StartAPI()
			return
		case "0":
			fmt.Println("Saindo do menu Projects...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
