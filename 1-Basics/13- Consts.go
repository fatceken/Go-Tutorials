package main

import "fmt"

const(
	Zero int =  iota * 2
	Two 
	Four
	Six 
	Eight
)

func main(){

	fmt.Println(Zero)
	fmt.Println(Two)
	fmt.Println(Eight)
	
}