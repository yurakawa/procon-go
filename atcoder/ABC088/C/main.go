package main

import "fmt"

func main() {
	var c [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j< 3; j++ {
			fmt.Scan(&c[i][j])
		}
	}

	var a [3]int
	var b [3]int
	for i := 0; i < 3; i++ {
		b[i] = c[0][i] - a[0]
	}

	for i := 0; i < 3; i++ {
		a[i] = c[i][0] - b[0]
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i]+b[j] != c[i][j] {
				fmt.Println("No")
				return

			}
		}
	}
	fmt.Println("Yes")
}
