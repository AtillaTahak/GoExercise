package main

import (
	"fmt"
	"sort"
)
var productName string

type Product struct {
	name       string
	price      float64
	stock      bool
	stockCount int
}

func (p *Product) Buy() {
	if p.stock && p.stockCount > 0 {
		p.stockCount--
		fmt.Printf("Purchased %s. New stock count: %d\n", p.name, p.stockCount)
		if p.stockCount <=0{
			p.stock = false;
		}
	} else {
		p.stock=false;
		fmt.Printf("Product %s is out of stock.\n", p.name)
	}
}

func searchProduct (products []Product ,name string){
	for _,product := range products{
		if product.name == name{
			fmt.Printf("Your Product %+v",product)
		}
	}
}
func showOnStockProduct(products []Product){
	for _,product := range products{
		if product.stock{
			fmt.Printf("%+v this product have stock and count of stock %v",product, product.stockCount)
		}
		fmt.Printf("%+v this product have not stock",product)
	}
}
func main() {
	products := []Product{
		{"Apple", 1.99, true, 20},
		{"Banana", 0.99, true, 10},
		{"Orange", 1.49, true, 5},
	}

	for chosen := true; chosen; {
		fmt.Println("1. Print All Products")
		fmt.Println("2. Buy Products with Name")
		fmt.Println("3. Search Products")
		fmt.Println("4. Show Just Have Stock Product")
		fmt.Println("5. Show Product with highest Product")
		fmt.Println("6  Show Product with lowest Product")
		fmt.Println("7. Exit")
		fmt.Println("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("1. Print All Products")
			for _, product := range products {
				fmt.Println(product)
			}
		case 2:
			fmt.Println("2. Buy Products with Name")
			fmt.Println("Enter the name of the product to buy: ")
			fmt.Scanln(&productName)
			for i := range products {
				if products[i].name == productName {
					products[i].Buy()
					break
				}
			}
		case 3:
			fmt.Println("3. Search Products")
			fmt.Println("Enter the name of the product to search: ")
			fmt.Scanln(&productName)
			searchProduct(products , productName)
		case 4:
			fmt.Println("4. Check Product Stock")
			showOnStockProduct(products)
		case 5:
			fmt.Println("5. Show Product with highest Product")
			sort.Slice(products, func(i, j int) bool {
				return products[i].price < products[j].price
			})
			fmt.Print(products)
		case 6:
			fmt.Println("5 Show Product with Lowest Product")
			sort.Slice(products, func(i, j int) bool {
				return products[i].price > products[j].price
			})
			fmt.Print(products)
		case 7:
			chosen = false
		default:
			fmt.Println("Invalid choice")
		}
	}
}
