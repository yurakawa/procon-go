package main

import "fmt"

func main() {
	var x1,x2,x3,x4,x5 int
	fmt.Scan(&x1,&x2,&x3,&x4,&x5)

	if x1 == 0 {
		fmt.Println(1)
	}
	if x2 == 0 {
		fmt.Println(2)
	}
	if x3 == 0 {
		fmt.Println(3)
	}
	if x4 == 0 {
		fmt.Println(4)
	}
	if x5 == 0 {
		fmt.Println(5)
	}
}

