package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	max := Max{v: 0, index: n + 1}
	max2 := Max{v: 0, index: n + 1}
	for i := 0; i < n; i++ {
		if max.v <= a[i] {
			max2 = max
			max = Max{v: a[i], index: i}
		} else if max2.v <= a[i] {
			max2 = Max{v: a[i], index: i}
		}
	}
	for i := 0; i < n; i++ {
		if i == max.index {
			fmt.Println(max2.v)
			continue
		}
		fmt.Println(max.v)
	}
}

type Max struct {
	v     int
	index int
}
