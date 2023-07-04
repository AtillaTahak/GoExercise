package main

type OrderFactory struct{}

func (o OrderFactory) createOrder(orderType string, quantity int, unitPrice int) Order{
	switch orderType{
	case "Pizza":
		return PizzaOrder{quantity, unitPrice}
	case "Burger":
		return BurgerOrder{quantity, unitPrice}
	default:
		return nil
	}
}