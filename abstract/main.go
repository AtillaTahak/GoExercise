package main

type Doer interface {
	Do()
}

type test struct {
}

func (t test) Do() {
	println("test")
}

func main() {
	var d Doer
	d = test{}
	d.Do()
}