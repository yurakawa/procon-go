package main

import "fmt"

func main() {
	var x,n  int
	fmt.Scan(&x,&n)

	p := make([]int, n )
	for i:= 0; i < n; i++ {
		fmt.Scan(&p[i])
	}
	if n == 0 {
		fmt.Println(x)
		return
	}

	for i := 0;i <n+100; i++ {
		lit :=  x - i
		big := x + i
		litOk := true
		bigOk := true
		for _, v := range p {
			if lit == v {
				litOk = false
			}
			if big == v {
				bigOk = false
			}
		}
		if litOk {
			fmt.Println(lit)
			return
		}
		if bigOk {
			fmt.Println(big)
			return
		}
	}
}

