package main

import (
	"fmt"
)

func main() {
	fmt.Print("Please enter a number: ")
	var number int
	fmt.Scanln(&number)
	fmt.Println("Factorial of ", number, " is: ", factorial(number))
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}