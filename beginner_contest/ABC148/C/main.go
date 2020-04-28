package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(lcm(a,b))

}

func gcd(a, b int) int {
	if a > b {
		return gcd(b, a)

	}
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
