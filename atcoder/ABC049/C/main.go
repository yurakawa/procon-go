package main

import (
	"fmt"
	"regexp"
)

// https://atcoder.jp/contests/abc049/tasks/arc065_a
// 白昼夢

func main() {
	var s string
	fmt.Scan(&s)

	divide := `^(dream|dreamer|erase|eraser)*$`

	if regexp.MustCompile(divide).Match([]byte(s)) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
