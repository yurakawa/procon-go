package main

import (
	"fmt"
)

func main() {
	var x, k, d int
	fmt.Scan(&x, &k, &d)

	if x == 0 {
		if k%2 == 0 {
			fmt.Println(0)
			return
		}
		fmt.Println(d)
		return
	}
	x = abs(x)
	ans := x % d
	a := x / d
	if a > k {
		// 往復に辿り着く前に終る
		fmt.Println(x - k*d)
		return
	} else if a <= k {
		// 往復する
		if (k-a)%2 == 0 {
			fmt.Println(ans)
			return
		} else if (k-a)%2 == 1 {
			if ans > 0 {
				ans -= d
			} else {
				ans += d
			}
			fmt.Println(abs(ans))
			return
		}
	}
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
