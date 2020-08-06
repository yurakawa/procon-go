package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	d := nextInt()

	cnt := 0

	for i := 0; i < n; i++ {
		x, y := nextInt(), nextInt()
		if d *d  >=  x * x + y *y{
			cnt++
		}
	}
	fmt.Println(cnt)

}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
