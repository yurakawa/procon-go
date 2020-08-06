package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	cnt := 0
	for {
		n = n - 2 * d - 1
		cnt++
		if n <= 0 {
			fmt.Println(cnt)
			return
		}
	}
}
