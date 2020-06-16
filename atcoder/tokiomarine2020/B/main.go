package main

import "fmt"

func main() {
	var a,v,b,w, t int
	fmt.Scan(&a, &v, &b, &w, &t)

	if w -v > 0 {
		fmt.Println("NO")
		return
	}
	if abs(b-a) <= (t* abs(w-v)) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
