package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Please select file")
	var fileName string
	fmt.Scanln(&fileName)
	file, err := os.ReadFile(fileName)
	if err != nil{
		fmt.Println("There some errors", err)
	}
	content := string(file[:])
	fmt.Println("This is your content", content)

}