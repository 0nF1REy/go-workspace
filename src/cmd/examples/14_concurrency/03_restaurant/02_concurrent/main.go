package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mesas := []string{"Mesa 1", "Mesa 2", "Mesa 3"}

	var wg sync.WaitGroup

	for _, mesa := range mesas {
		wg.Add(1)
		go atender(mesa, &wg)
	}

	wg.Wait()
}

func atender(mesa string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Atendendo", mesa)
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizado", mesa)
}