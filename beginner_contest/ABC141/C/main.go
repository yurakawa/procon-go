package main

import "fmt"

func main() {
	var n, k, q int
	fmt.Scan(&n, &k, &q)

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		scores[i] = k - q
	}
	corCnt:= map[int]int{}
	for i := 0; i < q ; i++ {
		var a int
		fmt.Scan(&a)
		corCnt[a-1]++
	}


	for k, v := range corCnt {
		scores[k] += v
	}
	for _, v := range scores {
		if v > 0 {
			fmt.Println("Yes")
			continue
		}
		fmt.Println("No")
	}
}

