package main

import (
	"fmt"
	"strings"
)


// マージソート
// https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_5_B

var infty = 1000000000 * 2

func merge (s []int, left, mid, right int) int {
	nLeft := mid - left
	nRight := right - mid

	pLeft := make([]int, nLeft+1)
	pRight := make([]int, nRight+1)

	for i := 0; i < nLeft; i++ {
		pLeft[i] = s[left+i]
	}
	for i := 0; i < nRight; i++ {
		pRight[i] = s[mid+i]
	}

	pLeft[nLeft] = infty
	pRight[nRight] = infty

	i,j := 0, 0
	for k := left; k < right; k++ {
		if pLeft[i] <= pRight[j] {
			s[k] = pLeft[i]
			i++
		} else {
			s[k] = pRight[j]
			j++
		}
	}

	return right -left
}

func mergeSort(s []int, left, right int) int {
	if left+1 >= right {
		return 0
	}
	mid := (left + right) / 2
	nCompareLeft := mergeSort(s, left, mid)
	nCompareRight := mergeSort(s, mid, right)
	nCompareMerge := merge(s, left, mid, right)

	return nCompareLeft + nCompareRight + nCompareMerge
}

func main(){
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	//	fmt.Println(s)
	nCompare := mergeSort(s, 0, n)
	fmt.Println(stringifyArray(s))
	fmt.Println(nCompare)
}

func stringifyArray(arr []int) string {
	return strings.TrimRight(fmt.Sprintf("%+v", arr)[1:], "]")
}
