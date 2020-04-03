package main

// https://atcoder.jp/contests/abs/tasks/abc081_a

import "fmt"

func main() {
	var str string

	fmt.Scan(&str)
	var c int
	for _,s := range str {
		if s == '1' {
			c++
		}
	}
	fmt.Println(c)
}
