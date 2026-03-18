package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	results := make(chan int, 100)

	for i := range 100 {
		wg.Add(1)

		go func(port int) {
			defer wg.Done()

			addr := fmt.Sprintf("scanme.nmap.org:%d", port)
			conn, err := net.Dial("tcp", addr)

			if err == nil {
				fmt.Printf("Porta %d aberta em scanme.nmap.org\n", port)
				conn.Close()
				results <- 1
			} else {
				fmt.Printf("Porta %d fechada em scanme.nmap.org\n", port)
				results <- 0
			}
		}(i)
	}

	wg.Wait()
	close(results)

	total := 0
	for r := range results {
		total += r
	}

	fmt.Println("==============================")
	fmt.Println("Scan concluído")
	fmt.Printf("Total de portas abertas: %d\n", total)
	fmt.Println("==============================")
}
