package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Taking input from user
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	// Calculator using switch-case
	switch operator {
	case "+":
		fmt.Printf("Result: %.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("Result: %.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("Result: %.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Invalid operator. Please use one of +, -, *, /")
	}
}
