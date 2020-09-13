package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s []byte
	fmt.Scan(&s) // 1905

	ans := "NA"
	a, _ := strconv.Atoi(string(s[:2])) // 19
	b, _ := strconv.Atoi(string(s[2:])) //05
	if a >= 1 && a <= 12 {
		ans = "MMYY"
	}
	if b >= 1 && b <= 12 {
		if ans == "MMYY" {
			fmt.Println("AMBIGUOUS")
			return
		}
		fmt.Println("YYMM")
		return
	}
	fmt.Println(ans)
}
