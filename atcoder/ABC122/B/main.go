package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	max := 0
	for i := 0; i < len(s); i++ {
		cnt := 0
		for j := i; j < len(s); j++ {
			if s[j] == 'A' ||
				s[j] == 'C' ||
				s[j] == 'G' ||
				s[j] == 'T' {
				cnt++
				continue

			}
			break
		}
		if max < cnt {
			max = cnt
		}
	}
	fmt.Println(max)
}
