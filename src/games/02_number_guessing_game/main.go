/*
	Jogo de adivinhação:
	- Usuário escolhe a dificuldade
	- Sistema gera um número aleatório dentro de um intervalo
	- Usuário tem número limitado de tentativas
	- Feedback: muito alto / muito baixo
	- Vitória ao acertar, derrota ao acabar as tentativas
*/

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Escolha o nível de dificuldade (Difícil - 1, Médio - 2, Fácil - 3)")

	var difficulty int
	fmt.Scan(&difficulty)

	var min_range int
	var max_range int
	var total_life int

	switch difficulty {
	case 1: // Difícil
		min_range = 0
		max_range = 500
		total_life = 5

	case 2: // Médio
		min_range = 0
		max_range = 100
		total_life = 7

	case 3: // Fácil
		min_range = 0
		max_range = 50
		total_life = 10

	default:
		fmt.Println("Nível de dificuldade inexistente.")
		return
	}

	fmt.Printf("Tente adivinhar um número entre %d e %d\n", min_range, max_range-1)

	var user_input int

	system_number := rand.IntN(max_range-min_range) + min_range

	for {
		if total_life <= 0 {
			fmt.Println(`
          _______
       .-"       "-.
      /             \
     |   X       X   |
     |      ___      |
     |     (___)     |
      \             /
       \  \_____/  /
        "-._____.-"
    `)
			fmt.Println("VOCÊ PERDEU!")
			fmt.Printf("O número correto era: %d\n", system_number)
			break
		}

		fmt.Printf("Tentativas restantes: %d\n", total_life)
		fmt.Print("Digite o seu palpite -> ")
		fmt.Scanln(&user_input)

		// Validações
		if user_input < min_range {
			fmt.Println("Palpite abaixo do intervalo permitido.")
			continue
		}

		if user_input >= max_range {
			fmt.Println("Palpite acima do intervalo permitido.")
			continue
		}

		if user_input == system_number {
			fmt.Println(`
        .     '     ,
          _________
       _ /_|_____|_\ _
         '. \   / .'
           '.\ /.'
             '.'
    `)
			fmt.Println("VOCÊ GANHOU!")
			break
		} else if user_input < system_number {
			fmt.Println("Palpite muito baixo.")
		} else {
			fmt.Println("Palpite muito alto.")
		}

		total_life--
	}
}
