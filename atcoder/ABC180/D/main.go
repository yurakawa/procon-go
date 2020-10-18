package main

import (
	"fmt"
)

func main() {
	var x, y, a, b int
	fmt.Scan(&x, &y, &a, &b)

	exp := 0
	for x <= (x+b)/a && x < y/a {
		exp++
		x *= a
	}
	fmt.Println(exp + ((y - 1 - x) / b))
}
