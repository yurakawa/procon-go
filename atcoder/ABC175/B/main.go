package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&l[i])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if i == j || i == k || j == k {
					continue
				}

				if l[i]+l[j] > l[k] &&
					l[j]+l[k] > l[i] &&
					l[k]+l[i] > l[j] {
					if l[i] != l[j] && l[j] != l[k] && l[k] != l[i] {
						///fmt.Println(i,j,k)
						//fmt.Println(l[i],l[j],l[k])
						cnt++

					}
				}
			}
		}
	}
	fmt.Println(cnt)
}
