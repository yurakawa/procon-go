package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	b := make([]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		b[i], _ = strconv.Atoi(sc.Text())
	}

	ans := 0
	for i := 0; i < n; i++ {
		p := 0
		for j := 0; j < m; j++ {
			sc.Scan()
			a, _ := strconv.Atoi(sc.Text())
			p += a * b[j]
		}
		if p+c > 0 {
			ans++
		}
	}
	fmt.Println(ans)

}

