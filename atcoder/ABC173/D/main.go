package main

import (
	"fmt"
)


func main() {
	var n int
	fmt.Scan(&n)
	// pf := make([]int, n)
	// mod := int(1e9 + 7)
	sum := 0
	for i := 1; i <= n; i++ {
		ans = enumDivisions(i)
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
func enumDivisions(n int) int {
	var res []int
	sum := 0
	for i := 1; i * i <= n; i++ {
		if n % i == 0 {
			res = append(res, i)

			// 重複しないならばiの相方である n/i もappend
			if n/i != i {
				res = append(res, n/i)
			}
			sum += res[i] * i
		}
	}
	return sum
}
