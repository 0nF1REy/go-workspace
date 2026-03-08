package functions

import "fmt"

func hello() {
    fmt.Println("Olá! Esta é uma função simples.")
}

func sum(a int, b int) int {
    return a + b
}

func swap(a, b string) (string, string) {
    return b, a
}

func average(numbers ...float64) float64 {
    total := 0.0
    for _, n := range numbers {
        total += n
    }
    return total / float64(len(numbers))
}

var greet = func(name string) string {
    return "Olá, " + name + "!"
}

func Execute() {
    fmt.Println("=== Demonstração de Funções em Go ===")

    hello()

    result := sum(5, 7)
    fmt.Println("Soma de 5 + 7 =", result)

    a, b := swap("primeiro", "segundo")
    fmt.Println("Swap:", a, b)

    avg := average(4.0, 8.0, 10.0)
    fmt.Println("Média:", avg)

    message := greet("Alan")
    fmt.Println(message)
}
