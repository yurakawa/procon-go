package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := 1; i <= n ;i++ {
		if len(strconv.Itoa(i)) % 2 == 1 {
			ans++
		}
	}

	fmt.Println(ans)
}
