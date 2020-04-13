package main

import "fmt"

// C - Maximum Volume
// https://atcoder.jp/contests/abc159/tasks/abc159_c

// 長方形の等周問題,相加相乗平均の話
// 直方体の中で最も体積が最大になるのは立方体
func main() {
	var l float64
	fmt.Scan(&l)

	x, y, z:= l / 3, l / 3, l / 3
	ans := x * y * z
	fmt.Println(ans)
}
