package main

// https://atcoder.jp/contests/abs/tasks/abc085_b

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() int{
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	n := nextLine()

	m := make(map[int]struct{})
	var num int
	for i := 0; i < n; i++ {
		num = nextLine()
		if _, ok := m[num]; !ok {
			m[num] = struct{}{}
		}
	}
	fmt.Println(len(m))
}
