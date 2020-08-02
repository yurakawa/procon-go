package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	aSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)

		aSum[i+1] = a + aSum[i]
	}
	store := make(map[int]int, n+1)

	cnt := 0
	for i := 0; i < n+1; i++ {
		cnt+= store[aSum[i]]
		store[aSum[i]]++

	}
	fmt.Println(cnt)

}
