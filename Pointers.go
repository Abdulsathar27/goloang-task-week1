package main

import "fmt"


func modifyValue(num *int) {
	*num = *num * 2 
}

func main() {
	value := 10
	fmt.Println("Before modification:", value)

	modifyValue(&value) 

	fmt.Println("After modification:", value)
}
