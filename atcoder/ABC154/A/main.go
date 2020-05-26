package main

import "fmt"

func main() {
	var s, t string
	var a, b  int
	var u string
	fmt.Scan(&s, &t, &a, &b, &u)

	if s == u {
		a--
	}
	if t == u {
		b--
	}
	fmt.Println(a, b)
}
