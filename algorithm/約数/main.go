package main

import (
	"fmt"
	"sort"
)

// ABC052 C
// https://atcoder.jp/contests/abc052/tasks/arc067_a

func main() {
	var n int
	fmt.Scan(&n)
	pf := make([]int, 1010)
	mod := int(1e9 + 7)
	for i := 2; i <= n; i++ {
		primeFactor(i, pf)
	}
	ans := 1
	for j := 1; j < len(pf); j++ {
		if pf[j] >= 1 {
			ans = (ans * (pf[j] + 1)) % mod
		}
	}
	fmt.Println(ans)
}
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
