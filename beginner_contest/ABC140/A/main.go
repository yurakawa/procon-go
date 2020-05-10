package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(pow(n, 3))
}

func pow(a, b int) (ret int) {
	ret = a
	// 10^2 = 100 は10に10を1回掛けること
	// つまり初期値を含めると上限b-1未満
	for i := 0; i < b-1; i++ {
		ret *= a
	}
	return
}
