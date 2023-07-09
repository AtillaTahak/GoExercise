package main

type Player struct {
	Name string
	Age int
}

func calculateValues(a int, b int) int {
	return a + b + 1
}

func main() {
	calculateValues(1, 2)
}