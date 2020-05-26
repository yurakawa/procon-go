package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	ss:=split(s)
	switch ss[len(ss)-1]  {
	case "2","4","5","7","9":
		fmt.Println("hon")
	case "0", "1", "6", "8":
		fmt.Println("pon")
	case "3":
		fmt.Println("bon")

	}
}


func split(s string) []string {
	ret := make([]string, len([]rune(s)))
	for i, v := range []rune(s) {
		ret[i] = string(v)
	}
	return ret
}
