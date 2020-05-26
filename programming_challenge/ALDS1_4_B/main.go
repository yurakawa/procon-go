package main

import "fmt"

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_4_A

func main (){

	var n int
	fmt.Scan(&n)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	var q int
	fmt.Scan(&q)

	ans := make([]int, 0)
	for i := 0; i < q; i++ {
		var t int
		fmt.Scan(&t)
		for _, v := range s {
			if v == t {
				ans = append(ans, t)
			}
		}
	}
	fmt.Println(len(ans))
}

