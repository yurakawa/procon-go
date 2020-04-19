package main

import "fmt"

func main() {

	var n, m int
	fmt.Scan(&n, &m)

	var p int
	var s string
	q := make(map[int]bool, n)
	sumPenaQ := make(map[int]int, n)
	var sumPena int
	var sumAC int
	for i := 0; i < m; i++ {
		fmt.Scan(&p, &s)

		if q[p] == true {
			continue
		}

		if s == "AC" {
			sumAC++
			sumPena += sumPenaQ[p]
			q[p] = true
		} else { // "WA"
			sumPenaQ[p]++
		}
	}
	fmt.Println(sumAC, sumPena)
}
