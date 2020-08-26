package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 貰うDP

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	as := make([]bool, n)
	for i := 0; i < m; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		as[num] = true
	}

	dp := make([]int, n+1)
	mod := 1000000007

	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		// i - 1段目が安全なら
		if !as[i-1] {
			dp[i] += dp[i-1]
		}

		// i - 2段目が安全なら
		if !as[i-2] {
			dp[i] += dp[i-2]
		}
		dp[i] %= mod
	}
	fmt.Println(dp[n])

}
