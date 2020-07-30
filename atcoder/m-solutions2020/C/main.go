package main

import (
	"fmt"
)

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := k; i < n;i++ {
		if a[i] > a[i-k]{
			fmt.Println("Yes")
			continue
		}
		fmt.Println("No")
	}
}
