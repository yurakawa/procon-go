package main

import "fmt"

func main() {
	var a,b,c,d int
	fmt.Scan(&a,&b,&c, &d)

	for i:= 0; i < 100; i++ {
		c = c-b
		a =a-d
		if c <=0 {
			fmt.Println("Yes")
			return

		}
		if a <=0 {

			fmt.Println("No")
			return
		}
	}

}

