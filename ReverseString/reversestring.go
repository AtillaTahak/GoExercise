package main


import (
	"fmt"
)

func main() {
	fmt.Print("Please enter a string: ")
	var str string
	fmt.Scanln(&str)
	fmt.Println("Reverse of ", str, " is: ", reverse(str))
}
func reverse(str string) string {
	if len(str) <= 1 {
		return str
	}
	return reverse(str[1:]) + string(str[0])
}