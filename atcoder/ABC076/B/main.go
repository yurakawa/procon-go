package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	num := 1
	for i := 0; i < n; i++  {
		a := num * 2
		b := num + k
		if a > b {
			num = b
			continue
		}
		num = a
	}
	fmt.Println(num)
}
