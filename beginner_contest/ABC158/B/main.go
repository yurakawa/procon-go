package main

import (
	"fmt"
	"math"
)

func main() {
	var n, a, b int

	fmt.Scan(&n, &a, &b)

	c := a+b
	fmt.Println(a * (n / c) + min(n % c, a))
}


func min(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}
