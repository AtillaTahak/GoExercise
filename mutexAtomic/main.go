package main

import "fmt"

type State struct {
	count int
}

func main() {
	state := State{}
	for i := 0; i < 10; i++ {
		state.count++
	}

	fmt.Println(state)
}
