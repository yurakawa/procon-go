package main

import (
	"fmt"
	"math"
)

// 貰うDP(動的計画法)
/*
func main() {
	var n int
	fmt.Scan(&n)

	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}

	dp := make([]int, n+10)
	dp[1] = abs(h[1] - h[0])
	for i := 2; i < n; i++ {
		dp[i] = min(
			dp[i-2]+abs(h[i]-h[i-2]),
			dp[i-1]+abs(h[i]-h[i-1]),
		)
	}
	fmt.Println(dp[n-1])
}
*/
func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
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

const (
	INF = 1000000000000000000
)

// 配るDP
func main() {

	var n int
	fmt.Scan(&n)

	h := make([]int, n+10)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}

	dp := make([]int, n+10)
	for i := 0; i < n; i++ {
		dp[i] = INF
	}
	dp[1] = abs(h[1] - h[0])
	if n > 2 {
		dp[2] = abs(h[0] - h[2])
	}

	for i := 0; i < n; i++ {
		if i < n-1 {
			dp[i+1] = min(dp[i+1], dp[i]+abs(h[i]-h[i+1]))
		}
		if i < n-2 {
			dp[i+2] = min(dp[i+2], dp[i]+abs(h[i]-h[i+2]))
		}

	}
	fmt.Println(dp[n-1])
}
