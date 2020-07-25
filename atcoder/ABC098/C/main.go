package main

import (
	"fmt"
	"math"
)

// 累積和
// 事前に各indexまでのWとEの合計を保持して、
// 最後に各indexのpreWとsufEを足し合わせて最小値を取得する

func main() {
	var n int
	var s string

	fmt.Scan(&n, &s)

	cumW := make([]int, n+1)
	cumE := make([]int, n+1)

	for i := 0; i < n ; i++ {
		if string(s[i]) == "W" {
			cumW[i+1] = cumW[i] + 1
		} else {
			cumW[i+1] = cumW[i]
		}
	}

	for i := 0; i < n ; i++ {
		if string(s[i]) == "E" {
			cumE[i+1] = cumE[i] + 1
		} else {
			cumE[i+1] = cumE[i]
		}
	}

	answer:= make([]int, n)
	for i := 0 ; i< n; i++ {
		answer[i] = cumW[i] + (cumE[n] - cumE[i+1])
	}

	fmt.Println(min(answer...))
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
