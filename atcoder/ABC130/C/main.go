package main

import "fmt"

// 図形

func main() {
	var w, h, x, y float64
	fmt.Scan(&w, &h, &x, &y)
	// どっちが小さいか
	num := 0
	if w == x*2 && h == y*2 {
		num = 1
	}
	fmt.Println(w*h/2, num)

}
