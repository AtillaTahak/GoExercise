package main

import (
	"fmt"
)

func main() {
	fmt.Print("Please enter a number: ")
	var number int
	fmt.Scanln(&number)
	fmt.Println("Fibonacci of ", number, " is: ", fibonacci(number))
}

func fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
}