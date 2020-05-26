package main

import "fmt"

func main() {
	var a,b int
	fmt.Scan(&a, &b)

	n := 0
	for a*n-(n-1) < b {
		n++
	}

	fmt.Println(n)
}
