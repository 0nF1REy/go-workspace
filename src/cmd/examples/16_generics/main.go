package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

type stringer interface {
	String() string
}

type MyString string

func somaInteiros(m map[string]int64) int64 {
	var soma int64
	for _, v := range m {
		soma += v
	}
	return soma
}

func somaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func somaGenerica[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func soma[T comparable](number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}
	return number2
}

func max[T constraints.Ordered](number1 T, number2 T) T {
	if number1 > number2 {
		return number1
	}
	return number2
}

func (m MyString) String() string {
	return string(m)
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

func main() {
	var x, y, z MyNumber
	x = 1
	y = 2
	z = 3

	// SOMA
	fmt.Println("Soma Inteiros:", somaInteiros(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println("Soma Float:", somaFloat(map[string]float64{"a": 1.1, "b": 22.2, "c": 3.2}))
	fmt.Println("Soma Genérica Float:", somaGenerica(map[string]float64{"a": 1.1, "b": 22.2, "c": 3.2}))
	fmt.Println("Soma Genérica MyNumber:", somaGenerica(map[string]MyNumber{"a": x, "b": y, "c": z}))

	// MAX
	fmt.Println("Max entre x e y (MyNumber):", max(x, y))
	fmt.Println("Max entre 10 e 20 (int):", max(10, 20))
	fmt.Println("Max entre 3.14 e 2.71 (float64):", max(3.14, 2.71))

	// Soma comparável
	fmt.Println("Soma comparável x e y:", soma(x, y))
	fmt.Println("Soma comparável 5 e 5:", soma(5, 5))

	// Concatenação de strings
	fmt.Println("Concatenação MyString:", concat([]MyString{"a", "b", "c"}))
}
