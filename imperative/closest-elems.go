package main

import (
	"fmt"
	"sort"
)

func main() {
	// test := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}
	// fmt.Println(findClosestElements(test, 3, 5))

	test2 := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}
	fmt.Println(findClosestElements(test2, 3, -5))

	// test3 := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}
	// fmt.Println(findClosestElements(test3, 3, 15))

	// test4 := []int{1, 2, 3, 4, 5}
	// fmt.Println(findClosestElements(test4, 4, 3))
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func findBracket(arr []int, x int) (lo, hi int) {
	for i, val := range arr {
		if val > x {
			return i - 1, i
		}
	}
	return len(arr) - 2, len(arr) - 1
}

func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) == 0 {
		return nil
	}

	closestElems := make([]int, 0, k)

	lptr, rptr := findBracket(arr, x)

	for len(closestElems) != k {
		if lptr < 0 && rptr >= len(arr) {
			break
		}

		if lptr < 0 {
			closestElems = append(closestElems, arr[rptr])
			rptr++
			continue
		}
		if rptr >= len(arr) {
			closestElems = append(closestElems, arr[lptr])
			lptr--
			continue
		}

		if abs(x, arr[lptr]) > abs(arr[rptr], x) {
			closestElems = append(closestElems, arr[rptr])
			rptr++
		} else {
			closestElems = append(closestElems, arr[lptr])
			lptr--
		}
	}

	sort.Ints(closestElems)
	return closestElems
}
