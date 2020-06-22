package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(strings.ToLower(toAlphabet(n)))
}

func toAlphabet(self int) string {
	n := self % 26
	if n == 0 {
		n = 26
	}
	s := string(n+64)

	if self == n {
		return s
	}
	return toAlphabet((self - n ) / 26) + s
}
