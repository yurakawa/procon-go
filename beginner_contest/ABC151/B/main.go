package main

import "fmt"

func main() {

	var n, k, m int
	fmt.Scan(&n, &k, &m)

	var a int
	var sum int
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a)
		sum += a
	}
	ans := m * n - sum

	if ans < 0 {
		fmt.Println("0")
		return

	}
	if k < ans{
		fmt.Println("-1")
		return
	}
	fmt.Println(ans)
}
