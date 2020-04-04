package main

import (
	"fmt"
	"strings"
)

// https://atcoder.jp/contests/abs/tasks/arc065_a
// 白昼夢


func main() {
	var s string
	fmt.Scan(&s)

	s = reverse(s)

	for {
		s = strings.Replace(s, "resare", "", -1)
		s = strings.Replace(s, "esare", "", -1)
		s = strings.Replace(s, "remaerd", "", -1)
		s = strings.Replace(s, "maerd", "", -1)
	}
	if len(s) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
