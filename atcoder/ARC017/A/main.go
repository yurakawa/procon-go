package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
