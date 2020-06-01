package main

import "fmt"

func main() {
	a := []int{1, 1, 2, 2, 2, 4, 5, 5, 6, 8, 8, 8, 10, 15}
	i := lower_bound(a, 3)
	fmt.Printf("a[%d] = %d\n", i, a[i])

	i = lower_bound(a, 2)
	fmt.Printf("a[%d] = %d\n", i, a[i])
}


// ソートされた範囲に対するアルゴリズム
// 指定した値Value以上の最初の要素の位置をイテレータで返す

func lower_bound(arr []int, value int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < value {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
