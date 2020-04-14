package main

import "fmt"

// https://atcoder.jp/contests/abc158/tasks/abc158_a

func main() {
	var s string
	fmt.Scan(&s)


	slice := split(s)
	if slice[0] == slice[1] && slice[0]== slice[2] {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
func split(s string) []string {
	ret := make([]string, len([]rune(s)))
	for i, v := range []rune(s) {
		ret[i] = string(v)
	}
	return ret
}

