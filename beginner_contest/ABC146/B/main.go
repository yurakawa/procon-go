package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n,&s)

	for _, c := range s {
		x := c + int32(n)
		if x > 90 {
			x -=26
		}
		fmt.Printf("%c",x)
	}
}

