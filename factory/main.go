package main


func main(){
	orderFactory := OrderFactory{}
	pizzaOrder := orderFactory.createOrder("Pizza", 2, 10)
	pizzaOrder.printReceipt()
	burgerOrder := orderFactory.createOrder("Burger", 3, 5)
	burgerOrder.printReceipt()
}