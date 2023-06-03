package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Please write some number")
	var number int
	fmt.Scanln(&number)
	fmt.Println("Number of Digits : ", numberDigits(number))
}
func numberDigits(number int) int{
	numberString := strconv.Itoa(number)
	 return len(numberString)
}