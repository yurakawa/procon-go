package main

import "fmt"

// 全探索で行けるんかい
func main() {
	var x int
	fmt.Scan(&x)

	for i := -500; i <= 500; i++ {
		for j := -500; j <= 500; j++ {
			if five(i) - five(j) == x {
				fmt.Println(i,j)
				return
			}
		}
	}
}

func five(l int) int{
	return l*l*l*l*l
}

