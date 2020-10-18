package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	ns := enumDivisions(n)
	for _, v := range ns {
		fmt.Println(v)
	}
}

// 約数列挙
func enumDivisions(n int) []int {
	var res []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
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
