package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

type animal struct {
	name string
}

func renameAnimalByValue(a animal) {
	a.name = "valteri"
}

func renameAnimalByRef(a *animal) {
	a.name = "lewis"
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	//structs are value types in go

	cat := animal{name: "Garfield"}
	fmt.Println(cat)
	renameAnimalByValue(cat)
	fmt.Println(cat)
	renameAnimalByRef(&cat)
	fmt.Println(cat)

}
