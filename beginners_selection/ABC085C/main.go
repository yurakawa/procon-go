package main

import (
	"bufio"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abs/tasks/abc085_c

import (
	"fmt"
)

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
	N, Y := nextInt(), nextInt()

	var z int
	for x := 0; x <= N; x++ {
		for y := 0; y <= (N - x); y++ {
			z = N - (x + y)
			if (x * 10000 + y * 5000 + z * 1000) == Y {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")

}

