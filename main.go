package main

import (
	"fmt"

	lom "github.com/samber/lo/mutable"
	lop "github.com/samber/lo/parallel"

	"github.com/samber/lo"
)

func main() {
	fmt.Println("=== Lo Library Examples ===")

	// https://pkg.go.dev/github.com/samber/lo#readme-spec
	// 1: map, filter, reduce

	numbers := []int{1, 2, 3, 4, 5}
	// numbers_2 := new([]int) returns the pointer
	// numbers_3 := make([]int, 5)

	numbersSquared := lo.Map(numbers, func(x int, i int) int {
		return x * x
	})

	numbersFiltered := lo.Filter(numbersSquared, func(item int, index int) bool { return item > 2 })

	sum := lo.Reduce(numbersFiltered, func(agg int, item int, index int) int { return agg + item }, 61)

	fmt.Printf("numbers: %v, numbers squared: %v, numbers filtered: %v, reduced sum with 61: %v\n", numbers, numbersSquared, numbersFiltered, sum)

	// parallel map, unique map
	// and mutable map :/

	type footballer struct {
		name string
		age  int
	}

	ts := []footballer{{name: "oulai", age: 19}, {name: "zubi", age: 31}, {name: "batagov", age: 24}, {name: "batagov", age: 24}, {name: "batagov", age: 24}}
	fmt.Println("initial ts: ", ts)

	tsUniq := lo.UniqMap(ts, func(item footballer, index int) string {
		return item.name
	})

	fmt.Println("ts uniq: ", tsUniq)

	tsAfterAYear := lop.Map(ts, func(item footballer, index int) footballer {
		item.age += 1
		return item
	})

	fmt.Println("ts after a year (everyone grows paralelly :p): ", tsAfterAYear)

	lom.Map(ts, func(item footballer) footballer {
		item.age += 1
		return item
	})

	fmt.Println("ts after mutation: ", ts)
	// lets try reverse too
	toRev := []int{1, 2, 3}

	fmt.Println("initial list: ", toRev)
	lom.Reverse(toRev)

	fmt.Println("reversed list: ", toRev)
	lom.Shuffle(toRev)

	fmt.Println("shuffled list: ", toRev)
	// tomorrow: other functions of lop, partition by group by etc.

	someNums := []int{1, 2, 3, 4, 6, 7, 9}
	partitioned := lop.PartitionBy(someNums, func(item int) int { return item % 3 })
	fmt.Printf("partitioned %v\n", partitioned)

	timesLooksWeird := lop.Times(5, func(i int) int { return i * 2 })
	fmt.Printf("times %v\n", timesLooksWeird)

	type order struct {
		accountId int
		orderId   int
	}

	orders := []order{{accountId: 1, orderId: 1},
		{accountId: 1, orderId: 2},
		{accountId: 2, orderId: 3}}

	result := lop.GroupBy(orders, func(item order) int {
		return item.accountId
	})

	fmt.Printf("orders grouped by: %v\n", result)
	// orders grouped by: map[1:[{1 1} {1 2}] 2:[{2 3}]]

}
