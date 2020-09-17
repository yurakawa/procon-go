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

	h := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		h[i], _ = strconv.Atoi(sc.Text())
	}

	ans := 1
	max := h[0]
	for i := 1; i < n; i++ {
		if h[i] < max {
			continue
		}
		ans++
		if max < h[i] {
			max = h[i]
		}
	}

	fmt.Println(ans)

}
