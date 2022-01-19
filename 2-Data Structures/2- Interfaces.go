package main

import (
	"fmt"
)

type MyInt int

func main() {

	typeAssertion()

}

func typeAssertion() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt) // like 'as' keyword in .net  ( for ex : myHandler as HttpHandler)
	fmt.Println(i2 + 1)
}
