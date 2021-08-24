package main

import "fmt"

func main() {

	x := 5
	y := 6

	result := x < y
	result2 := x > y
	result3 := x <= y
	result4 := x >= y
	result5 := x != y
	result6 := x == y

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)

	left := true
	right := false

	result7 := left && right
	fmt.Println(result7)

	result8 := left || right
	fmt.Println(result8)

}
