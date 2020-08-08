package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	pfs := primeFactor(n)
	cnt := 0
	for _, v := range pfs {
		s := 0
		for i := v; s < i; i -= s {
			cnt++
			s++
		}
	}
	fmt.Println(cnt)
}

// 素因数分解
func primeFactor(n int) map[int]int {
	ret := make(map[int]int)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			ret[i]++
		}
	}
	if n > 1 {
		ret[n]++
	}
	return ret
}
