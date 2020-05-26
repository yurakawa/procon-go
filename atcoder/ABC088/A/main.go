package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n, &a)

	n = n % 500
	if n <= a {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")


}
