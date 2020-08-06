package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		fmt.Scan(&a[i])

	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if a[i] >= b[i] {
			cnt += b[i]
		} else {
			sub := b[i] - a[i]
			cnt += a[i]
			if a[i+1] > sub {
				cnt += sub
				a[i+1] -= sub
			} else {
				cnt += a[i+1]
				a[i+1] = 0
			}
		}
	}
	fmt.Println(cnt)
}
