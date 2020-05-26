package main

import (
	"fmt"
	"math"
)

// 全部列挙する

func main() {
	var n int
	fmt.Scan(&n)

	positions := make([][]int, n)
	for i :=0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		positions[i] = []int{x,y}
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			diffX := float64(positions[i][0] - positions[j][0])
			diffY := float64(positions[i][1] - positions[j][1])
			ans += math.Sqrt(diffX*diffX + diffY*diffY)
		}
	}
	fmt.Println(ans/float64(n))
}
