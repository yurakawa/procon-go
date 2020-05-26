package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scan(&n)

	as := make([]float64, 200000)
	for i := 0; i < n; i++ {
		fmt.Scan(&as[i])
	}

	var maxv, minv float64
	maxv = -2000000000
	minv = as[0]

	for i := 1; i < n; i++ {
		maxv = math.Max(maxv, as[i]-minv)
		minv = math.Min(minv, as[i])
	}
	fmt.Println(int(maxv))
}

