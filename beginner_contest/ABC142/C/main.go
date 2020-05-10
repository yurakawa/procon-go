package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	ans := make([]int, n)
	for i := 0; i < n;i++ {
		var a int
		fmt.Scan(&a)
		ans[a-1] = i+1
	}
	fmt.Println(strings.Join(arrayIntToArrayString(ans), " "))
}

func arrayIntToArrayString(arr []int) (r []string) {
	for _, v := range arr {
		r = append(r, fmt.Sprintf("%d",v ))
	}
	return
}
