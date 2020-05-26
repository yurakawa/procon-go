package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	ans := n * (n-1) / 2 + m * (m-1) / 2
	fmt.Println(ans)
}
