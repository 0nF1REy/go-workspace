package menu

import (
	"fmt"
	"os"
	"os/exec"
)

func Examples(enable bool) {
	if !enable {
		return
	}

	examples := GetExamples()

	for {
		fmt.Println("\n=== CLI Interativo [EXAMPLES] ===")
		for i, ex := range examples {
			fmt.Printf("%d - %s\n", i+1, ex.Name)
		}
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			continue
		}

		if input == "0" {
			fmt.Println("Saindo do menu Examples...")
			return
		}

		idx := -1
		for i := range examples {
			if input == fmt.Sprintf("%d", i+1) {
				idx = i
				break
			}
		}

		if idx >= 0 && idx < len(examples) {
			cmd := exec.Command("go", "run", examples[idx].MainPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			fmt.Printf("\n--- Rodando exemplo: %s ---\n", examples[idx].Name)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Erro ao rodar exemplo:", err)
			}
		} else {
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
