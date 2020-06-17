package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		m[a]++
	}


	var t []int
	for _, v := range m {
		t = append(t, v)
	}
	sort.Ints(t)

	wcount := 0
	for i := 0; i < len(t); i++ {
		if len(t) > k {
			k++
			wcount += t[i]
			continue
		}
		fmt.Println(wcount)
		return
	}
}
