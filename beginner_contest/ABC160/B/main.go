package main

import "fmt"

func main(){
	var x int
	fmt.Scan(&x)

	u := x / 500 * 1000
	u2 := x % 500 / 5 * 5
	fmt.Println(u+u2)
}
