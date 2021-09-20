package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func alternativePlus(a, b int) int {
	return a + b
}

func iterativeFact(a int) int {

	if a == 0 || a == 1 {
		return 1
	}

	var result = 1

	for i := 2; i <= a; i++ {
		result = result * i
	}

	return result
}

func recursiveFact(a int) int {

	if a == 0 || a == 1 {
		return 1
	}

	return a * recursiveFact(a-1)
}

func multipleReturn(a int) (int, string, bool) {
	return a, "helloWorld", true
}

func dynamicParameters(numbers ...int) int { //like args in .net
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {

	result := plus(1, 2)
	fmt.Println("result =", result)

	result2 := alternativePlus(1, 2)
	fmt.Println("result2 =", result2)

	result3 := iterativeFact(5)
	fmt.Println("result3 =", result3)

	result4 := recursiveFact(5)
	fmt.Println("result4 =", result4)

	multipleResult1, multipleResult2, multipleResult3 := multipleReturn(1)
	fmt.Println("multipleResult1 =", multipleResult1)
	fmt.Println("multipleResult2 =", multipleResult2)
	fmt.Println("multipleResult3 =", multipleResult3)

	_, multipleResult4, _ := multipleReturn(1)
	fmt.Println("multipleResult4 =", multipleResult4)

	result5 := dynamicParameters(1, 2, 3)
	fmt.Println("result5 =", result5)

	result6 := dynamicParameters(1, 2, 3, 4)
	fmt.Println("result6 =", result6)

	nums := []int{5, 6, 7}
	result7 := dynamicParameters(nums...)
	fmt.Println("result7 =", result7)

}
