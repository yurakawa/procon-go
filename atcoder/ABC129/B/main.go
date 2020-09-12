package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	w := make([]int, n)
	sum := make([]int, n+1)
	sum[0] = 0
	for i := 0; i < n; i++ {
		sc.Scan()
		w[i], _ = strconv.Atoi(sc.Text())
		sum[i+1] = sum[i] + w[i]
	}

	min := 10000

	for i := 0; i < n; i++ {
		g1 := sum[i]
		g2 := sum[n] - sum[i]
		if min > abs(g2-g1) {
			min = abs(g2 - g1)
		}
	}
	fmt.Println(min)
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
