package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := true
	for k, v := range split(s) {
		i := k + 1
		if i % 2 == 0 {
			if v == "R"  {
				ans = false
			}
		} else {
			if v == "L"  {
				ans = false
			}
		}
	}
	if ans {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")

}

func split(s string) []string {
	ret := make([]string, len([]rune(s)))
	for i, v := range []rune(s) {
		ret[i] = string(v)
	}
	return ret
}
