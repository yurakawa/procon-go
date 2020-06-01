package main

import "fmt"

// 係数ソート(数え上げソート)
// counting sort
// https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_6_A

func main() {
	var n int
	fmt.Scan(&n)

	c := make([]int, int(1e4)+1)

	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		c[v]++
	}

	cnt := 0

	for i := 0; i < len(c); i++ {
		if c[i] == 0 {

		}

	}


}

