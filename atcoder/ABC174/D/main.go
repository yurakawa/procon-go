package main

import "fmt"

func main() {
	var n int
	var c []byte
	fmt.Scan(&n, &c)

	jWhited := n -1
	cnt := 0
	for i := 0; i<n; i++ {

		if c[i] == 'W' {
			for j := jWhited; i < j; j-- {
				if c[j] == 'R' {
					c[j] = 'W'
					c[i] = 'R'
					cnt++
					break
				}
			}
		}
	}
	fmt.Println(cnt)
}
