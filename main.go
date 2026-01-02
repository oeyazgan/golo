package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	fmt.Println("=== Lo Library Examples ===\n")

	// 1: map, filter, reduce

	numbers := []int{1, 2, 3, 4, 5}
	// numbers_2 := new([]int) returns the pointer
	// numbers_3 := make([]int, 5)

	numbersSquared := lo.Map(numbers, func(x int, i int) int {
		return x * x
	})

	numbersFiltered := lo.Filter(numbersSquared, func(item int, index int) bool { return item > 2 })

	sum := lo.Reduce(numbersFiltered, func(agg int, item int, index int) int { return agg + item }, 61)

	fmt.Printf("numbers: %v, numbers squared: %v, numbers filtered: %v, reduced sum with 61: %v", numbers, numbersSquared, numbersFiltered, sum)
}
