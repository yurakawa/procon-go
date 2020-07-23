package main

import "fmt"

func main() {
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)

	cnt := 0
	for i := 0; i <= 100; i++ {
		if a <= i  && b > i {
			if c <= i  && d > i {
				cnt++
			}

		}
	}
	fmt.Println(cnt)
}
