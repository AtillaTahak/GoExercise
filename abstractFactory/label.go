package main

import (
	"fmt"
)

type Label interface{
	printLabel()
}

type MacLabel struct{
}
func (m MacLabel) printLabel(){
	fmt.Println("Mac Label")
}

type WindowsLabel struct{
}
func (w WindowsLabel) printLabel(){
	fmt.Println("Windows Label")
}