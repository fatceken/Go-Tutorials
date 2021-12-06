package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go f1(ch)
	go f2(ch)

	for {
		select {

		case n := <-ch:
			fmt.Println(n)

		default:

		}
	}

}

func f1(c chan int) {
	c <- 1
}

func f2(c chan int) {
	c <- 2
}
