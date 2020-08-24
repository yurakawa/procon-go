package main

import (
	"fmt"
	"math"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	x := make([][]int, n)
	for i := 0; i < n; i++ {
		x[i] = make([]int, d)
		for j := 0; j < d; j++ {
			fmt.Scan(&x[i][j])
		}
	}
	ans := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sq := 0
			for k := 0; k < d; k++ {
				diff := x[i][k] - x[j][k]
				sq += diff * diff
			}

			s := int(math.Sqrt(float64(sq) + 0.5))
			if s*s == sq {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
