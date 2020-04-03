package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abs/tasks/abc087_b

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
	a, b, c, x := nextLine(), nextLine(), nextLine(), nextLine()
	var cnt int

	for i := 0 ; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if x == i*500+j*100+k*50 {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
