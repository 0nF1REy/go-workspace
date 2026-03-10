package main

import (
	"fmt"

	"github.com/0nF1REy/go-workspace/cmd/examples/11_auth/app"
)

func main() {

	for {
		fmt.Println("\n=== Sistema de Login ===")
		fmt.Println("1 - Criar usuário")
		fmt.Println("2 - Login")
		fmt.Println("3 - Sair")

		var option int
		fmt.Print("Escolha: ")
		fmt.Scanln(&option)

		switch option {

		case 1:

			var username string
			var password string

			fmt.Print("Novo usuário: ")
			fmt.Scanln(&username)

			fmt.Print("Senha: ")
			fmt.Scanln(&password)

			err := app.Register(username, password)

			if err != nil {
				fmt.Println("Erro:", err)
				continue
			}

			fmt.Println("Usuário criado com sucesso.")

			hash, exists := app.GetEncryptedPassword(username)
			if exists {
				fmt.Println("Senha criptografada armazenada:", hash)
			}

		case 2:

			var username string
			var password string

			fmt.Print("Usuário: ")
			fmt.Scanln(&username)

			fmt.Print("Senha: ")
			fmt.Scanln(&password)

			auth := app.Authenticate(username, password)

			if !auth {
				fmt.Println("Usuário ou senha inválidos. Tente novamente.")
				continue
			}

			session := app.NewSession(username)

			fmt.Printf("\nUsuário %s logado com sucesso!\n", session.Username)
			fmt.Println("Encerrando...")
			return

		case 3:

			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida")

		}
	}
}
