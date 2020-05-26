package main

import (
	"fmt"
)

func main() {
	var a,b,c,k int
	fmt.Scan(&a,&b,&c,&k)

	if k <= a {
		fmt.Println(k)
		return
	}
	if k <= a+ b {
		fmt.Println(a)
		return
	}

	fmt.Println(a-(k-a-b))
	return
}


