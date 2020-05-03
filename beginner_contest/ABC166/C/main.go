package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	hs := make([]int, n)
	for i := 0 ; i < n;i++ {
		var h int
		fmt.Scan(&h)
		hs[i] = h

	}

	ans := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		ans[i] = true
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		if hs[a-1]> hs[b-1] {
			ans[b-1] = false
		} else if hs[a-1] < hs[b-1] {
			ans[a-1] = false
		} else {
			ans[a-1] = false
			ans[b-1] = false
		}
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if ans[i] == true {
			cnt++
		}
	}

	fmt.Println(cnt)
}
