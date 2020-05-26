package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	fmt.Println(len(strconv.FormatInt(int64(n), k)))
}
