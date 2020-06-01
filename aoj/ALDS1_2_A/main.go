package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}


	a, sw := bubbleSort(a, n)
	fmt.Println(a)
	fmt.Println(sw)


}

func bubbleSort(a []int, n int) ([]int ,int) {
	sw := 0
	flag := true

	for i := 0; flag; i++ {
		flag = false
		for j := n-1 ; j >= i + 1; j-- {
			if a[j] < a[j-1] {
				a[j-1], a[j] = a[j], a[j-1]
				flag =true
				sw++
			}
		}
	}
	return a, sw

}
