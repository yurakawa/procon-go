package main

import "fmt"

// 全探索
// https://onlinejudge.u-aizu.ac.jp/problems/ALDS1_5_A

var n int
var a [50]int

// i 階層
// m 与えられた数を引いていって0となったら狙った数値を作成できたことを示す
func solve(i, m int) bool{
	if m == 0 {return true}
	if i >= n {return false}
	return solve(i + 1, m) || solve(i + 1,  m - a[i])
}

func main() {
	//var n int
	fmt.Scan(&n)
	// a := make([]int,n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var m int
		fmt.Scan(&m)
		if solve(0, m) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
