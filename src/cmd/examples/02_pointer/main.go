package main

import (
	"fmt"
)

type Car struct {
	Name string
}

func (c *Car) drive() {
	c.Name = "Mazda RX-7 FD3S"
	fmt.Println(c.Name, "andou!")
}

// Função que dobra o valor de um inteiro via ponteiro
func doubleValue(num *int) {
	*num = *num * 2
}

func setValue(num *int) {
	*num = 200
}

func setValueViaPointer() {
	value := 10
	setValue(&value)
	fmt.Println(value)
}

func basicPointer() {
	x := 10
	px := &x

	fmt.Println("Endereço de x:", px)
	fmt.Println("Valor de x via ponteiro:", *px)
}

func changeValueViaPointer() {
	x := 10
	px := &x

	*px = 20
	fmt.Println("Novo valor de x após alteração via ponteiro:", x)
}

func pointerFunction() {
	y := 5
	fmt.Println("\nAntes da função, y =", y)

	doubleValue(&y)

	fmt.Println("Depois da função, y =", y)
}

func pointerDifferentTypes() {
	str := "Alan"
	pStr := &str

	fmt.Println("\nEndereço da string:", pStr)
	fmt.Println("Valor da string via ponteiro:", *pStr)
}

func carPointerMethod() {
	car := Car{
		Name: "Toyota AE86 Trueno",
	}

	car.drive()
	fmt.Println(car.Name)
}

func main() {
	basicPointer()
	changeValueViaPointer()
	pointerFunction()
	pointerDifferentTypes()
	setValueViaPointer()
	carPointerMethod()
}
