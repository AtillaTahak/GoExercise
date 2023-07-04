package main

import (
	"fmt"
)

type Button interface{
	printButton()
}

type MacButton struct{
}

func (m MacButton) printButton(){
	fmt.Println("Mac Button")
}

type WindowsButton struct{
}

func (w WindowsButton) printButton(){
	fmt.Println("Windows Button")
}
