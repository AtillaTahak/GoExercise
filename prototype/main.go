package main

import "fmt"

// Prototip struct'ı tanımlama
type ToyPrototype struct {
	Name  string
	Color string
}

// Prototipin davranışlarını gerçekleştiren metotlar
func (tp *ToyPrototype) Play() {
	fmt.Printf("Playing with %s toy\n", tp.Name)
}

// Prototipin bir kopyasını döndüren Clone fonksiyonu
func (tp *ToyPrototype) Clone() *ToyPrototype {
	return &ToyPrototype{
		Name:  tp.Name,
		Color: tp.Color,
	}
}

func main() {
	// Prototipi oluşturma
	prototype := &ToyPrototype{
		Name:  "Car",
		Color: "Red",
	}

	// Prototipin kopyasını oluşturma
	toy1 := prototype.Clone()
	toy1.Name = "Truck"

	// Kopyalanan prototipin davranışlarını kullanma
	toy1.Play()

	// Başka bir kopya oluşturma
	toy2 := prototype.Clone()
	toy2.Name = "Robot"

	// Diğer kopyalanan prototipin davranışlarını kullanma
	toy2.Play()
}
