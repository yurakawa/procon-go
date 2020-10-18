package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		x[i], _ = strconv.Atoi(sc.Text())
	}

	ans1 := 0
	ans2 := 0
	ans3 := 0
	for i := 0; i < n; i++ {
		ans1 += abs(x[i])
		ans2 += x[i] * x[i]
		ans3 = max(ans3, abs(x[i]))
	}
	fmt.Println(ans1)
	fmt.Println(math.Sqrt(float64(ans2)))
	fmt.Println(ans3)

}

// http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
func Dist(source, dest []float64) float64 {
	val := 0.0
	for i, _ := range source {
		val += math.Pow(source[i]-dest[i], 2)
	}
	return math.Sqrt(val)
}

func max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if res < num {
			res = num
		}
	}
	return res
}
