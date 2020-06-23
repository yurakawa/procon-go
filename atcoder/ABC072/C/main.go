package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	counter := make(map[int]int,n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)

		counter[a]++
		counter[a-1]++
		counter[a+1]++
	}

	max := 0
	for _, v := range counter {
		if max < v {
			max = v
		}
	}
	fmt.Println(max)
}

