package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var ans string
	for i := 0; i < len(s); i++{
		ans += "x"
	}
	fmt.Println(ans)
}
