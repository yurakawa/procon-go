package main

import (
	"fmt"
	"regexp"
)

// https://atcoder.jp/contests/abs/tasks/arc065_a
// 白昼夢

func main() {
	var s string
	divide := `^(dream|dreamer|erase|eraser)*$`

	fmt.Scanf("%s", &s)

	if regexp.MustCompile(divide).Match([]byte(s)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
