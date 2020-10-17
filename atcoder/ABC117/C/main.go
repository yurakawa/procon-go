package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	l := make([]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		l[i], _ = strconv.Atoi(sc.Text())

	}

	sort.Ints(l)
	dis := make([]int, m-1)
	for i := 0; i < m-1; i++ {
		dis[i] = l[i+1] - l[i]
	}
	sort.Ints(dis)
	ans := 0
	for i := 0; i < m-n; i++ {
		ans += dis[i]
	}
	fmt.Println(ans)

}
