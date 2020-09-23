package main

import "fmt"

func main() {
	collection1 := []int{3, 2, 5, -4, 8, 11}
	collection2 := []int{3, 11, -4, 2, 8, 5}
	fmt.Println(foo(7, collection2))
	fmt.Println(foo(7, collection1))
}

func foo(input int, numbers []int) (int, []int) {
	var result []int
	var nextIndex int

	for idx, number := range numbers {
		nextIndex = idx + 1

		if nextIndex >= len(numbers) {
			break
		}
		nextNumber := numbers[nextIndex]
		if number+nextNumber == input {
			result = append(result, number)
			result = append(result, nextNumber)
		}
	}
	return input, result
}
