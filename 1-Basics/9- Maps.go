//also known as dictionary

package main

import "fmt"

func main() {

	sampleMap := make(map[string]int)

	sampleMap["key1"] = 7
	sampleMap["key2"] = 13

	fmt.Println("map:", sampleMap)

	value1 := sampleMap["key1"]
	fmt.Println("value1: ", value1)

	fmt.Println("length:", len(sampleMap))
	
	delete(sampleMap, "key2")
	fmt.Println("map:", sampleMap)
	
	value2, isInMap := sampleMap["key2"] // The optional second return value when getting a value from a map indicates if the key was present in the map
	fmt.Println("value2:", value2)
	fmt.Println("isInMap:", isInMap)
	
	_, isInMap2 := sampleMap["key2"] // No need the value itself, so ignore it with the blank identifier _
	fmt.Println("isInMap:", isInMap2)
	
	anotherMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", anotherMap)
}
