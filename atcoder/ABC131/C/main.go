package main

import (
	"fmt"
)

func main() {

	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	ans := f(b, c, d) - f(a-1, c, d)
	fmt.Println(ans)

}

func f(x, c, d int) int {
	res := x
	res -= x / c
	res -= x / d
	res += x / lcm(c, d)
	return res
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
