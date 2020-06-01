package main

import "fmt"

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_4_B
// 二分探索
// 整列されてるなら

var a [100000]int
var n int

func binarySearch(key int) bool {
	left := 0
	right := n
	mid := 0

	for left < right {
		mid = (left + right) / 2
		if key == a[mid]{
			return true
		}
		if key > a[mid] {
			left = mid + 1
		} else if key < a[mid] {
			right = mid
		}
	}
	return false
}

func main (){
	fmt.Scan(&n)

	for i :=0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var q int
	fmt.Scan(&q)

	sum := 0
	for i := 0; i < q; i++ {
		var k int
		fmt.Scan(&k)
		if binarySearch(k) {
			sum++
		}
	}
	fmt.Println(sum)
}

