package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)


	for a := 0 ; a <= n ; a++ {
		for b := 0 ; a + b <= n ; b++ {
			for c := 0 ; a + b +c <= n ; c++ {
				// if y < 10000 * c{
				// 	break
				// }

				// if a + b + c  > n {
				// 	break
				// }
				if a + b + c == n {
					if y == 10000 * c + 5000* b + 1000* a {
						fmt.Println(c, b, a)
						return
					}
				}
			}
		}
	}

	fmt.Println("-1 -1 -1")
}

