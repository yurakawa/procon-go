package main

import (
	"fmt"
	"strconv"
)
// 二分探索

func main (){
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	l := 0
	r := int(1e9) + 1
	for r-l > 1 {
		c := (r + l) / 2
		b := a*c + b*len(strconv.Itoa(c))
		if b <= x {
			l = c
		} else {
			r = c
		}
	}
	fmt.Println(l)
}
