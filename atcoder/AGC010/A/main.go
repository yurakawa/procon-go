package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)


	sum := 0
	for i := 0; i <n ; i++ {
		var a int
		fmt.Scan(&a)
		sum += a
	}

	if  sum % 2 == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
