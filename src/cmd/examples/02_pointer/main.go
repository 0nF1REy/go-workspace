package main

import (
    "fmt"
)

func main() {
    x := 10
    px := &x

    fmt.Println("Endereço de x:", px)
    fmt.Println("Valor de x via ponteiro:", *px)
}
