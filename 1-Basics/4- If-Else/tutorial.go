package main

import "fmt"

func main() {

	age := 12

	if age >= 18 {
		fmt.Println("Can drive")
	} else {
		fmt.Println("Can not drive")
	}

	score := 100

	if score <= 70 {
		fmt.Println("Drink never")
		} else if score > 70 && score < 100 {
			fmt.Println("Drink rarely")
		} else {	
			fmt.Println("Drink always")
		}
	}


