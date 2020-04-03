package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//https://atcoder.jp/contests/abs/tasks/abc083_b
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
	n, a, b := nextInt(), nextInt(), nextInt()

	var sum int
	for i := 1; i <= n; i++ {
		if a <= sumDigit(i)  && sumDigit(i) <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}

func sumDigit(n int) (sum int){
	for {
		if n <= 0 {
			break
		}
		sum += n % 10
		n = n / 10
	}
	return
}
