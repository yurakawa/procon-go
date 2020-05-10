package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var cnt int
	for i := 0; i < n; i++ {
		var h int
		fmt.Scan(&h)
		if h >= k {
			cnt++
		}
	}
	fmt.Println(cnt)

}
