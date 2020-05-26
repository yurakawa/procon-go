package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	hs := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&hs[i])
	}

	max := 0

	cnt := 0
	prev := hs[0]
	for i := 1; i < n; i++ {
		if hs[i] <= prev {
			cnt++
			prev = hs[i]
			if max < cnt {
				max = cnt
			}
			continue
		}
		cnt = 0
		prev = hs[i]
	}
	fmt.Println(max)
}
