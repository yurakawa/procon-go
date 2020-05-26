package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var a int
	var sum int
	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		sum += a
	}

	ans := n-sum
	if ans < 0 {
		fmt.Println("-1")
		return
	}
	fmt.Println(ans)
}
