package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	cnt := 0
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		if a[i] % 2 == 1 && i % 2 == 1 {
			cnt++
		}
	}

	fmt.Println(cnt)

}

