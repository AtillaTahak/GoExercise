package main

import "fmt"

// Sipariş yapısı
type Order struct {
	TableNumber int
	Meal        string
	Drink       string
	Dessert     string
}

// Sipariş yapısını oluşturmak için kullanılacak builder arayüzü
type OrderBuilder interface {
	SetTableNumber(int) OrderBuilder
	AddMeal(string) OrderBuilder
	AddDrink(string) OrderBuilder
	AddDessert(string) OrderBuilder
	Build() Order
}

// Concrete Builder
type ConcreteOrderBuilder struct {
	order Order
}

func (cob *ConcreteOrderBuilder) SetTableNumber(tableNumber int) OrderBuilder {
	cob.order.TableNumber = tableNumber
	return cob
}

func (cob *ConcreteOrderBuilder) AddMeal(meal string) OrderBuilder {
	cob.order.Meal = meal
	return cob
}

func (cob *ConcreteOrderBuilder) AddDrink(drink string) OrderBuilder {
	cob.order.Drink = drink
	return cob
}

func (cob *ConcreteOrderBuilder) AddDessert(dessert string) OrderBuilder {
	cob.order.Dessert = dessert
	return cob
}

func (cob *ConcreteOrderBuilder) Build() Order {
	return cob.order
}

// Director
type OrderDirector struct {
	builder OrderBuilder
}

func (od *OrderDirector) SetBuilder(builder OrderBuilder) {
	od.builder = builder
}

func (od *OrderDirector) BuildOrder(tableNumber int, meal, drink, dessert string) Order {
	return od.builder.
		SetTableNumber(tableNumber).
		AddMeal(meal).
		AddDrink(drink).
		AddDessert(dessert).
		Build()
}

// İstemci kodu
func main() {
	// Builder nesnesini oluştur
	builder := &ConcreteOrderBuilder{}

	// Director nesnesini oluştur ve Builder'ı atayarak bağımlılığı kur
	director := &OrderDirector{}
	director.SetBuilder(builder)

	// Director üzerinden siparişi oluştur
	order := director.BuildOrder(5, "Hamburger", "Coke", "Ice Cream")

	// Oluşturulan siparişi kullan
	fmt.Println(order)
}
