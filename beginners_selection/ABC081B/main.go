package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abs/tasks/abc081_b

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
	num := nextInt()
	as := make([]int, num)

	var cnt int

	for i := 0; i < num; i++ {

		n := nextInt()
		as = append(as, n)
		if n % 2 == 1 {
			fmt.Println(cnt)
			return
		}
	}
	for {
		for i, a := range as {
			if a % 2 == 1 {
				fmt.Println(cnt)
				return
			}
			as[i] = a / 2
		}
		cnt++
	}
}
