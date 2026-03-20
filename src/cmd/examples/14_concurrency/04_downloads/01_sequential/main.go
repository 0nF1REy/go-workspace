package main

import (
	"fmt"
	"time"
)

func main() {
	arquivos := []string{"arquivo1", "arquivo2", "arquivo3"}

	for _, a := range arquivos {
		baixar(a)
	}
}

func baixar(a string) {
	fmt.Println("Baixando", a)
	time.Sleep(2 * time.Second)
	fmt.Println("Concluído", a)
}
