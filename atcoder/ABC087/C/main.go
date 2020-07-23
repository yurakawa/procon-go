package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a1[i])
	}
	a2 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a2[i])
	}

	max := 0
	for i := 0; i < n; i++ {
		sum := 0
		right := n - i
		for left := 0 ; left < right; left++ {
			sum += a1[left]
		}
		right--
		for ;right < n; right++ {
			sum += a2[right]
		}
		if max < sum {
			max = sum
		}
	}
	fmt.Println(max)
}
