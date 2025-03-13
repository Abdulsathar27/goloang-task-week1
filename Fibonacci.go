package main

import "fmt"

func printFibonacciSeries(num int) {
	a := 0
	b := 1

	fmt.Printf("Fibonacci Series up to %d: %d %d ", num, a, b)

	for {
		c := a + b
		if c >= num {
			break
		}
		fmt.Printf("%d ", c)
		a = b
		b = c
	}
	fmt.Println() 
}

func main() {
	printFibonacciSeries(10)
	printFibonacciSeries(16)
	printFibonacciSeries(100)
}
