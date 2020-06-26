package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)


	max := max(a,b,c)
	sum := a+ b+c
	cnt := 0
	if 3 *max % 2 ==  sum %2 {
		for ; max*3 != sum; sum += 2 {
			cnt++
		}

	} else {
		for ; (max+1) * 3 != sum; sum += 2 {
			cnt++
		}
	}

	fmt.Println(cnt)
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
