package main

import (
    "fmt"
)

func hello() {
    fmt.Print("Hello" + " ")
}

func world() {
    fmt.Print("World")
}

func helloWorld() {
    hello()
    world()
    fmt.Println("!")
}

func main() {
    helloWorld()
}
