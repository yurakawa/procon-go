package main

import "fmt"

func main() {
	var n, q int
	var s string
	fmt.Scan(&n, &q, &s)

	cnt := make([]int, n)
	c := 0
	for i := 1; i < n;i++ {
		if string(s[i]) == "C" {
			if string(s[i-1]) == "A" {
				c++
			}
		}
		cnt[i] = c
	}

	for i := 0 ; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		fmt.Println(cnt[r-1] - cnt[l-1])
	}
}
