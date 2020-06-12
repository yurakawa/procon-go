package main

import (
	"fmt"
	"math"
	"sort"
)

// オイラー関数
// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=NTL_1_D&lang=ja

func main() {
	var n int
	fmt.Scan(&n)
	pf := make([]int,n)
	phi := n
	nReconstructed := 1
	primeFactor(n, pf)
	for i, p := range pf {
		if p > 0 {
			phi = phi * (i - 1) / i
			nReconstructed *= int(math.Pow(float64(i), float64(p)))
		}
		if nReconstructed >= n {
			break
		}
	}
	fmt.Println(phi)
}
// 素因数分解
func primeFactor(n int, pf []int) {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			pf[i]++
		}
	}
	if n > 1 {
		pf[n]++
	}
}

// 約数列挙
func enumDivisions(n int) []int {
	var res []int
	for i := 1; i * i <= n; i++ {
		if n % i == 0 {
			res = append(res, i)

			// 重複しないならばiの相方である n/i もappend
			if n/i != i {
				res = append(res, n/i)
			}
		}
	}
	// 小さい順に並び替える
	sort.Ints(res)
	return res
}
