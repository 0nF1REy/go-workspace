package main

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
	"github.com/gin-gonic/gin"
)

// menu recebe boolean para ativar ou não
func menu(enable bool) {
	if !enable {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n=== CLI Interativo Go Workspace ===")
		fmt.Println("1 - Tipos básicos")
		fmt.Println("2 - Estruturas de controle")
		fmt.Println("3 - Funções")
		fmt.Println("4 - Slices")
		fmt.Println("5 - Structs")
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
		case "0":
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}

// ginServer recebe boolean para ativar ou não
func ginServer(enable bool) {
	if !enable {
		return
	}

	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")
}

// Função central que decide o que vai rodar
func runApp(menuEnabled bool, ginEnabled bool) {
	menu(menuEnabled)
	ginServer(ginEnabled)
}

func main() {
	runApp(false, true)
}
