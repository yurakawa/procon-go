package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	if x == 1 {
		fmt.Println(0)

		return
	}
	fmt.Println(1)

}
