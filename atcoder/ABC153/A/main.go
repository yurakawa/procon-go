package main

import "fmt"

func main() {
	var h, a int
	fmt.Scan(&h, &a)

	var cnt int
	for {
		cnt++
		h -= a
		if h <= 0 {
			fmt.Println(cnt)
			return
		}
	}
}
