package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	as := make([]int, n + 1)
	for i := 0; i < n; i++ {
		fmt.Scan(&as[i])
	}

	cnt := 1
	prev := as[0]
	b :=  0 //   > 0,0 , 0 < の区分

	for i := 0; i < n; i++ {
		if b * (as[i]- prev) < 0 {
			cnt++
			b = 0
		} else if as[i] - prev != 0 {
			b = as[i] - prev
		}

		prev = as[i]
	}
	fmt.Println(cnt)
}
