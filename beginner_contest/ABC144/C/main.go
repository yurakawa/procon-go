package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := n
	for i := 1; i * i <= n; i++ {
		if n % i != 0 {
			continue
		}
		j := n / i
		if i+j - 2 < ans{
				ans = i+j-2
		}
	}
	fmt.Println(ans)
}
