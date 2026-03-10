package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
