package main

import (
	"fmt"
	"time"

	"github.com/0nF1REy/go-workspace/cmd/examples/10_user_welcome/app"
)

func main() {

	user := app.NewUser("Alan Ryan", "estudante")

	welcome := app.WelcomeMessage(user)

	fmt.Println(welcome)
	fmt.Printf("Data/Hora: %s\n", time.Now().Format("02/01/2006 - 15:04:05"))
}
