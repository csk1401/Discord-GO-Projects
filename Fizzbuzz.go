package main

import (
	"fmt"
)

func fizzBuzz(number int) {
	for i := 1; i <= number; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	fizzBuzz(number)
}
