package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := big.NewInt(1)

	upper := big.NewInt(1e18)
	as := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&as[i])
		if as[i]== 0 {
			fmt.Println(0)
			return
		}
	}
	for i, v := range  as {
		if i == 0 {
			sum = big.NewInt(v)
			continue
		}
		sum.Mul(sum, big.NewInt(v))
		if sum.Cmp(upper) >0 {
			fmt.Println("-1")
			return
		}
	}
	fmt.Println(sum.Int64())
}

