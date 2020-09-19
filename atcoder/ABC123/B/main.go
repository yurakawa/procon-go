package main

import (
	"fmt"
)

// abcde

func main() {
	a := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}

	sum := 0
	max := 0
	for i := 0; i < 5; i++ {
		sum += a[i]
		if v := a[i] % 10; v != 0 {
			p := 10 - v
			if p > max {
				max = p
			}
			sum += p
		}
	}
	fmt.Println(sum - max)

}
