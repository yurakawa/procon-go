package main

import (
	"fmt"
)

func main (){
	var n int
	fmt.Scan(&n)

	bs := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		var b int
		fmt.Scan(&b)
		bs[i] = b
	}
	as := make([]int, n)
	// ai= biとbi+1の小さい方
	for i := 0; i < n; i++ {
		if i == 0 {
			as[i] = bs[i]
			continue
		}
		if i == n - 1 {
			as[i] = bs[i-1]
			continue
		}
		if bs[i] > bs[i-1] {
			as[i] = bs[i-1]
		} else {
			as[i] = bs[i]
		}
	}
	fmt.Println(Sum(as...))
}

func Sum(arr ...int) (sum int) {
	for _, x := range arr {
		sum += x
	}
	return
}

