package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	as := make([]int, n)
	for i := 0 ; i < n;i++ {
		fmt.Scan(&as[i])
	}

	for ki := 0; ki < k; ki++ {
		b := make([]int, n+1)
		for x := 0; x < n; x++ {
			l := max(0, x-as[x])
			r := min(x-as[x]+1, n)
			b[l]++
			b[r]--
		}

		for x := 0; x <n; x++ {
			b[x+1] += b[x]
		}

		// b = b[:len(as)-1]
		if reflect.DeepEqual(as, b) {
			break
		}
		as = b
	}
	for i := 0; i < n; i ++ {
		fmt.Println(as[i])
	}
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
