//like foreach

package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}

	for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    
	sum := 0
    for _, num := range nums { 
        sum += num
    }
    fmt.Println("sum:", sum)

    

    keyValuePairs := map[string]string{"a": "apple", "b": "banana"}
    
	for key, value := range keyValuePairs {
        fmt.Printf("%s -> %s\n", key, value)
    }

    for k := range keyValuePairs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}