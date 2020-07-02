package main

import (
	"fmt"
	"sort"
)

func main (){
	var n, c, k int
	fmt.Scan(&n, &c, &k)

	t := make([]int, n)
	for i := 0; i <n ;i++ {
		fmt.Scan(&t[i])
	}
	sort.Ints(t)

	start := t[0]
	passenger:= 1
	ans := 1
	for i:= 1; i < n; i++ {
		if start + k < t[i] ||  passenger == c {
			ans++
			start = t[i]
			passenger = 1
			continue
		}
		passenger++
	}

	fmt.Println(ans)

}
