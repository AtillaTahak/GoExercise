package main

type Order interface{
	calculatePrice() int
	printReceipt()
}