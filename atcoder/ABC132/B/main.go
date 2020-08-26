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

	ps := make([]int, n)
	for i := range ps {
		sc.Scan()
		ps[i], _ = strconv.Atoi(sc.Text())
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		if ps[i-1] < ps[i] {
			if ps[i+1] > ps[i] {
				ans++
				continue
			}
		}
		if ps[i-1] > ps[i] {
			if ps[i+1] < ps[i] {
				ans++
			}
		}
	}
	fmt.Println(ans)

}
