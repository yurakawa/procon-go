package main

import "fmt"

// 累積和 x 素数

const lim = 1e5

func main() {
	var q int
	fmt.Scan(&q)


	primes := primes(lim)

	cnt := 0
	ans := make([]int, lim+1)
	for n := 1;  n <= lim; n++ {
		if primes[n] == 0 && (primes[(n+1)/2]) == 0 {
			cnt++
		}
		ans[n] = cnt
	}

	for i := 0 ;i < q ;i++ {
		var l, r int
		fmt.Scan(&l, &r)
		fmt.Println(ans[r] - ans[l-1])
	}
}
func primes(n int) []int {
	np := make([]int, n+1)
	np[0] = 1
	np[1] = 1
	for i := 2; i*i <= n; i++ {
		if np[i] == 0 {
			for j := i + i; j <= n; j += i {
				np[j] = 1
			}
		}
	}
	return np
}
