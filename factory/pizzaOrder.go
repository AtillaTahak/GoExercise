package main

import (
	"fmt"
)

type PizzaOrder struct{
	quantity int
	unitPrice int
}
func (p PizzaOrder) calculatePrice() int{
	return p.quantity * p.unitPrice
}
func (p PizzaOrder) printReceipt(){
	fmt.Printf("Pizza Order: %d\n", p.calculatePrice())
}