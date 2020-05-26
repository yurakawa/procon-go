package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	ans := a - b * 2
	if ans <= 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(ans)
}
