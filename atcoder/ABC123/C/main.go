package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	min := int(1e16)
	for i := 0; i < 5; i++ {
		var a int
		fmt.Scan(&a)
		if min > a {
			min = a
		}
	}
	p := 0
	if n%min != 0 {
		p = 1

	}

	fmt.Println(n/min + 4 + p)

}
