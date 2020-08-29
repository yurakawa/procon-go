package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// bit演算(全探索)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	s := make([][]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())

		for j := 0; j < k; j++ {
			sc.Scan()
			sj, _ := strconv.Atoi(sc.Text())
			s[i] = append(s[i], sj)
		}
	}

	p := make([]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		p[i], _ = strconv.Atoi(sc.Text())
	}

	var ans int
	for i := 0; i < 1<<n; i++ {
		on := make([]bool, n)
		for j := 0; j < n; j++ {
			// on[j] = (i>>j)%2 == 1
			on[j] = i>>j&1 == 1
		}
		var miss bool
		for j := 0; j < m; j++ {
			var num int
			for _, v := range s[j] {
				if on[v-1] {
					num++
				}
			}
			if num%2 != p[j] {
				miss = true
				break
			}
		}

		if !miss {
			ans++
		}
	}

	fmt.Println(ans)
}
