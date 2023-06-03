package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Please enter First Number: ")
	var numberOne string
	fmt.Scanln(&numberOne)
	fmt.Print("Please enter Second Number: ")
	var numberTwo string
	fmt.Scanln(&numberTwo)
	// how can i convert string to int
	 numberOneInt,_ := strconv.Atoi(numberOne)
	 numberTwoInt,_ := strconv.Atoi(numberTwo)
	fmt.Println("Sum of two numbers is: ", numberOneInt+numberTwoInt)

}
