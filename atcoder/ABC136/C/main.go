package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	h  := make([]int, n)
	for i := 0; i <n ; i++ {
		fmt.Scan(&h[i])
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			continue
		}
		if h[i] > h[i-1] {
			h[i]--
		}

		if h[i] < h[i-1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
