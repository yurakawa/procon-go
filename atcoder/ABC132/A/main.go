package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	m := make(map[string]int, len(s))
	for _, v := range s {
		m[string(v)]++
	}

	cnt := 0
	for _, v := range m {
		if v == 2 {
			cnt++
		}
	}

	if cnt == 2 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
