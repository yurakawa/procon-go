package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	var str string
	if a > b {
		for i := 0; i < a; i++ {
			str += strconv.Itoa(b)
		}
	} else {
		for i := 0; i < b; i++ {
			str += strconv.Itoa(a)
		}
	}
	fmt.Println(str)
}
