package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i], &c[i])
	}

	dp := make([][3]int, n+10)
	dp[0][0] = a[0]
	dp[0][1] = b[0]
	dp[0][2] = c[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][1]+a[i], dp[i-1][2]+a[i])
		dp[i][1] = max(dp[i-1][0]+b[i], dp[i-1][2]+b[i])
		dp[i][2] = max(dp[i-1][0]+c[i], dp[i-1][1]+c[i])
	}

	fmt.Println(max(dp[n-1][0], dp[n-1][1], dp[n-1][2]))

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

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
