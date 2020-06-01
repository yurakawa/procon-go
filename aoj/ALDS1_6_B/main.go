package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// パーティション
// https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_6_B


func partition(a []int ,p , r int) int {
	x := a[r]
	i := p - 1

	for j := p; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}


func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	index := partition(a, 0, n-1)
	buf := bytes.Buffer{}
	for i, e := range a {
		if i == index {
			buf.WriteString("[")
			buf.WriteString(strconv.Itoa(e))
			buf.WriteString("]")
		} else {
			buf.WriteString(strconv.Itoa(e))
		}
		buf.WriteString(" ")

	}
	fmt.Println(strings.TrimRight(buf.String(), " "))

}
