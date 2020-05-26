package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1; i < 1000000000; i++ {
		if a-i == -(b-i) {
			fmt.Println(i)
			return
		} else if  -(a-i) == b-i {
			fmt.Println(i)
		}
	}
	fmt.Println("IMPOSSIBLE")


}
