package main

import (
	"fmt"
	"net"
)

func main() {

	ports := 0

	for i := range 100 {
		addr := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", addr)

		if err == nil {
			fmt.Printf("Porta %d aberta em scanme.nmap.org\n", i)
			conn.Close()
			ports++
		} else {
			fmt.Printf("Porta %d fechada em scanme.nmap.org\n", i)
		}
	}

	fmt.Println("==============================")
	fmt.Println("Scan concluído")
	fmt.Printf("Total de portas abertas: %d\n", ports)
	fmt.Println("==============================")
}
