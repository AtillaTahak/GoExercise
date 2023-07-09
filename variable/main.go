package main

import (
	"fmt"
	"variable/util"
)

func main() {
	fmt.Println(util.Dive(10, 2))
	fmt.Println(util.Mines(10, 2))
	fmt.Println(util.Multiply(10, 2))
	fmt.Println(util.Sum(10, 2))
}