package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}


	kabu := 0
	cash := 1000

	for i := 0; i < n; i++ {
		if a[i] > a[i+1] {
			if kabu > 0 {
				// 売る
				cash += a[i] * kabu
				kabu = 0
			}
		} else {
			if kabu == 0 {
				// 買う
				kabu = cash / a[i]
				cash = cash % a[i]
			}
		}
	}
	fmt.Println(cash)

}
