package main

import "fmt"

func main() {
	var k, a, b int
	fmt.Scan(&k, &a, &b)

	if a <= b/k*k {
		fmt.Println("OK")
		return
	}
	fmt.Println("NG")

}
