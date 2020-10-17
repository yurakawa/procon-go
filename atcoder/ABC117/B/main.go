package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	l := make([]int, n)
	sum := 0
	max := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		l[i], _ = strconv.Atoi(sc.Text())
		sum += l[i]
		if max < l[i] {
			max = l[i]
		}
	}

	if max >= sum-max {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
