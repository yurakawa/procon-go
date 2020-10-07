package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sum := 0.0
	for i := 0; i < n; i++ {
		var x float64
		var u string
		fmt.Scan(&x, &u)
		if u == "BTC" {
			sum += x * 380000
		} else {
			sum += x
		}
	}
	fmt.Println(sum)
}
