package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Traveling
// https://atcoder.jp/contests/abs/tasks/arc089_a

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	N := nextInt()

	var nt, nx, ny int
	for i := 0; i < N; i++ {
		var t, x, y = nextInt(), nextInt(), nextInt()
		dt := t - nt
		diff := abs(nx-x) + abs(ny-y)
		if dt < diff || (dt > diff && (dt - diff)%2 != 0) {
			fmt.Println("No")
			return
		}
		nt, nx, ny = t, x, y
	}
	fmt.Println("Yes")
}

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	a, _ := strconv.Atoi(next())
	return a
}

// http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
