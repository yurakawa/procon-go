package main

import "fmt"

func main() {
	var a1, a2, a3 int
	fmt.Scan(&a1, &a2, &a3)

	if a1 + a2 + a3 > 21 {
		fmt.Println("bust")
	} else {
		fmt.Println("win")

	}
}
