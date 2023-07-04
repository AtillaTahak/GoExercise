package main

import (
	"fmt"
)

type BurgerOrder struct{
	quantity int
	unitPrice int
}
func (b BurgerOrder) calculatePrice() int{
	return b.quantity * b.unitPrice
}
func (b BurgerOrder) printReceipt(){
	fmt.Printf("Burger Order: %d\n", b.calculatePrice())
}
