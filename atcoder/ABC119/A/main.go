package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ABC119
func main() {
	var s string
	fmt.Scan(&s)

	ss := strings.Split(s, "/")
	d := strings.Join(ss, "")
	date, _ := strconv.Atoi(d)
	if date <= 20190430 {
		fmt.Println("Heisei")
	} else {
		fmt.Println("TBD")
	}
}

// func main() {
// 	var y, m, d int
// 	fmt.Scanf("%d/%d/%d\n", &y, &m, &d)
//
// 	if y < 2019 {
// 		fmt.Println("Heisei")
// 		return
// 	}
// 	if 2019 < y {
// 		fmt.Println("TBD")
// 		return
// 	}
// 	if m <= 4 {
// 		fmt.Println("Heisei")
// 		return
// 	}
// 	fmt.Println("TBD")
// }
