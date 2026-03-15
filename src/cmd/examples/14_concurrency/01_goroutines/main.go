package main

import (
	"fmt"
	"time"
)

func printCounter(label string) {
	for i := 0; i < 5; i++ {
		fmt.Println(label, i)
	}
}

func runBasicCounter(show bool) {

	if !show {
		return
	}

	printCounter("sem goroutine")
	go printCounter("com goroutine")

	fmt.Println("Hello A")
	fmt.Println("Hello B")

	time.Sleep(time.Second)
	fmt.Println("fim...")
}

func printCounterWithDelay(label string) {
	for i := 0; i < 5; i++ {
		fmt.Println(label, i)
		time.Sleep(time.Second)
	}
}

func runConcurrentCounters(show bool) {

	if !show {
		return
	}

	go printCounterWithDelay("a")
	go printCounterWithDelay("b")

	time.Sleep(10 * time.Second)
}

func main() {
	runBasicCounter(false)
	runConcurrentCounters(true)
}
