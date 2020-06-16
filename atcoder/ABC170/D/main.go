package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i:= 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	m := int(1e6+10)
	cnt := make([]int, m)
	for _,x := range a {
		if cnt[x] != 0 {
			cnt[x] = 2
			continue
		}
		for i := x; i <m; i += x {
			cnt[i]++
		}
	}
	ans := 0
	for _,x := range a {
		if cnt[x] == 1 {
			ans++
		}
	}
	fmt.Println(ans)
}

