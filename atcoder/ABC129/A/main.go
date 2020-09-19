package main

import "fmt"

func main() {
	var p, q, r int
	fmt.Scan(&p, &q, &r)

	fmt.Println(p + q + r - max(p, q, r))
}
func max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}
