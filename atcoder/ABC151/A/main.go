package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)

	fmt.Println(string([]rune(c)[0] + 1))
}
