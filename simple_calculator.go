package main

import (
	"fmt"
	"math"
)

func SimpleCalculator() {
	var choice int
	var num1, num2 float64

	fmt.Println("Simple Calculator")
	fmt.Println("Addition 1")
	fmt.Println("Subtraction 2")
	fmt.Println("Division 3")
	fmt.Println("Multiplication 4")
	fmt.Println("Power 5")

	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	fmt.Print("Enter two numbers: ")
	_, err = fmt.Scan(&num1, &num2)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	var result float64

	switch choice {
	case 1:
		result = num1 + num2
	case 2:
		result = num1 - num2
	case 3:
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			return
		}
		result = num1 / num2
	case 4:
		result = num1 * num2
	case 5:
		result = math.Pow(num1, num2)
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
