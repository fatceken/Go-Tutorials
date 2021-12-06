package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go producer(ch)

	fmt.Println("waiting")
	fmt.Println(<- ch, " consumed") 
	fmt.Println("Finished")
	
}

func producer(c chan string) {
	fmt.Println("start producing")
	time.Sleep(time.Second * 5)
	c <- "product"
	fmt.Println("produced")

}
