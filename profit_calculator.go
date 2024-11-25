package main

import (
	"fmt"
)

func ProfitCalculator() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter Revenue: ")
	_, err := fmt.Scan(&revenue)
	if err != nil {
		fmt.Println("Invalid input for revenue:", err)
		return
	}
	fmt.Print("Enter Expenses: ")
	_, err = fmt.Scan(&expenses)
	if err != nil {
		fmt.Println("Invalid input for expenses:", err)
		return
	}

	fmt.Print("Enter Tax Rate (in percentage): ")
	_, err = fmt.Scan(&taxRate)
	if err != nil {
		fmt.Println("Invalid input for tax rate:", err)
		return
	}

	EBT := revenue - expenses

	profit := EBT - (1 - taxRate/100)

	fmt.Printf("\nResults:\n")
	fmt.Printf("Revenue: %.2f\n", revenue)
	fmt.Printf("Expenses: %.2f\n", expenses)
	fmt.Printf("EBT (Earnings Before Taxes): %.2f\n", EBT)
	fmt.Printf("Profit After Tax: %.2f\n", profit)
}
