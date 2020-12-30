package main

import (
	"fmt"
)


func main() {
	var s []byte
	fmt.Scan(&s)


	ans := ""
	if string(s[len(s)-1]) == "s" {
		ans = "es"
	} else {
		ans = "s"
	}
	fmt.Println(string(s)+ans)



}

