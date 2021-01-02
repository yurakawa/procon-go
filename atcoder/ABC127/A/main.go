package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := 0
	if a >=  13 {
		ans = b
	} else if a >= 6 {
		ans = b/2
	}
	fmt.Println(ans)
}
