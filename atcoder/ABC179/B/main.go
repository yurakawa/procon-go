package main

import (
"bufio"
"fmt"
"os"
"strconv"
)


func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	ans := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		d1, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		d2, _ := strconv.Atoi(sc.Text())

		if d1== d2 {
			ans++
			if ans >= 3 {
				fmt.Println("Yes")
				return
			}
		} else {
			ans = 0
		}
	}

	fmt.Println("No")
}

