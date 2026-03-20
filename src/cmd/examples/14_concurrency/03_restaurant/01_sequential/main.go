package main

import (
	"fmt"
	"time"
)

func main() {
	mesas := []string{"Mesa 1", "Mesa 2", "Mesa 3"}

	for _, mesa := range mesas {
		atender(mesa)
	}
}

func atender(mesa string) {
	fmt.Println("Atendendo", mesa)
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizado", mesa)
}
