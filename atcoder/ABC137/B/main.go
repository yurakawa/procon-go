package main

import "fmt"

func main() {
	var k, x int
	fmt.Scan(&k, &x)

	min := x-k+1
	for i := 0; i < k + k -1; i ++{
		fmt.Print(min+i)
		fmt.Print(" ")
	}
}
