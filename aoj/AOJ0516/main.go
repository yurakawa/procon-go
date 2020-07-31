package main

import "fmt"

// 累積和
// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0516

func main() {

	for true {
		var n, k int
		fmt.Scan(&n, &k)

		if n == 0 && k == 0 {
			return
		}

		a := make([]int, n)
		ans := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
			if i == 0 {
				ans[i] = a[i]
			} else {
				ans[i] = ans[i-1] + a[i]
			}
		}
		max := 0
		for i := k; i < n; i++ {
			t := ans[i] - ans[i-k]
			if max < t {
				max = t
			}
		}
		fmt.Println(max)
	}
}
