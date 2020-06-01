package main

import "fmt"

func main() {
	var h1, m1, h2, m2, k int
	fmt.Scan(&h1, &m1, &h2, &m2, &k)
	ans := (h2*60 + m2) - (h1*60 + m1) - k
	fmt.Println(ans)
}
