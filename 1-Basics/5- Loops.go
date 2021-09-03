package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------")

	for i := 0; i < 10; i+=2 {
		fmt.Println(i)
	}
	
	for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            fmt.Println(i,j)
        }
	}
}
