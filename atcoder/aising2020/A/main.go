package main

import (
	"fmt"
)

func main() {
	var l, r, d int
	fmt.Scan(&l, &r, &d)

	cnt := 0
	for i := l; i <= r; i++ {
		if i %d == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}

// http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
