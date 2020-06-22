package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cnt := 0
	for i := 0; i < n; i++ {
		var p int
		fmt.Scan(&p)
		if p != i + 1 {
			cnt++
		}

		if cnt > 2 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
