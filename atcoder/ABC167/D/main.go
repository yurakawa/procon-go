package main

import (
	"fmt"
)

func main() {
	var n,k  int
	fmt.Scan(&n, &k)

	as := make([]int, n+1)
	for i:= 0; i < n; i++ {
		fmt.Scan(&as[i])
		as[i]--
	}

	ss := make([]int, 0)
	m := make(map[int]int)
	ss = append(ss, 0)
	m[0] = 0
	var cnt int
	for cnt < k {
		s := ss[len(ss)-1]
		if p, ok := m[as[s]]; ok {
			cycle := cnt - p + 1
			offset := (k - p) % (cycle)
			fmt.Println(ss[p+offset] + 1)
			return
		}
		ss = append(ss, as[s])
		cnt++
		m[as[s]] = cnt
	}
	fmt.Println(ss[len(ss)-1] + 1)
}


