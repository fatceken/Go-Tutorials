package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for _, val := range []string{"a", "b", "c"} {
		wg.Add(1)
		go func(val string) {
			fmt.Println(val)
			wg.Done()
		}(val)
	}
	wg.Wait()
}
