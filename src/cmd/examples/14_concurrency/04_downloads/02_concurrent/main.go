package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	arquivos := []string{"arquivo1", "arquivo2", "arquivo3"}

	var wg sync.WaitGroup

	for _, a := range arquivos {
		wg.Add(1)
		go baixar(a, &wg)
	}

	wg.Wait()
}

func baixar(a string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Baixando", a)
	time.Sleep(2 * time.Second)
	fmt.Println("Concluído", a)
}
