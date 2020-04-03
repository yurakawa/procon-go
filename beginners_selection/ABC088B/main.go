package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// https://atcoder.jp/contests/abs/tasks/abc088_b

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {

	sc.Split(bufio.ScanWords)
	n := nextInt()
	as := make([]int, n)

	for i := 1; i <= n; i++ {
		a := nextInt()
		as = append(as, a)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(as)))

	var bob, alice int

	for i, a := range as {
		if i % 2 == 0 {
			bob += a
		} else {
			alice += a
		}
	}
	fmt.Println(bob - alice)
}
