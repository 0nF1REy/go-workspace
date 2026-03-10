package app

import (
	"fmt"
	"strings"
)

func WelcomeMessage(user User) string {

	name := strings.TrimSpace(user.Name)
	role := strings.ToUpper(strings.TrimSpace(user.Role))

	return fmt.Sprintf("Bem-vindo, %s\nPerfil: %s", name, role)
}
