package main

import "fmt"

type Carro struct {
	Name string
}

func (c Carro) andar() {
	fmt.Println(c.Name, "andou!")
}

func sum(a, b int) int {
	return a + b
}

func sumAll(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}

func main() {

	carro := Carro {
		Name: "Toyota AE86 Trueno",
	}

	carro.andar()

	result := sumAll(4, 4, 1, 1) + sum(10, 10)
	fmt.Println(result)
}
