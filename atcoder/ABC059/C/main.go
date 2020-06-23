package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(min(solve(a, 0), solve(a, 1)))
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

func solve(a []int, ra int) int {
	var s, ans int

	for i, v := range a {
		if i%2 == ra {
			if s+v >= 0 {
				ans += abs(-1 - s - v)
				s = -1
				continue
			}
		} else {
			if s+v <= 0 {
				ans += abs(1 - s - v)
				s = 1
				continue
			}
		}
		s += v
	}
	return ans
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
