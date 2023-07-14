package main

import (
	"fmt"
	"time"
)

func main() {
	msgch := make(chan string,128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)

	for {
		msg, ok := <-msgch
		if !ok {
			break
		}
		fmt.Println(msg)
	}

/* 	for msg := range msgch {
		fmt.Println(msg)
	}
 */
}

func fetchResource(n int) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Resource %d", n)
}