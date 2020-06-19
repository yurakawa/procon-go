package main

import (
	"fmt"
	"math"
)

func main() {
	var a,b,c,x,y int
	fmt.Scan(&a,&b,&c,&x,&y)


	m := a * x + b * y
	for i := 1; i < x + y + 1; i++ {
		n := 2*c * i + max(0, x - i) *a +max(0, y -i) * b
		m = min(n, m)
	}
	fmt.Println(m)
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
func max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
