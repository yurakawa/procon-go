package main

import (
	"fmt"
)

// bit全探索

func main(){
	var h, w, k int
	fmt.Scan(&h, &w, &k )
	js := make([][]byte, h)
	for i := range js {
		fmt.Scan(&js[i])
	}

	var ans int

	for i := 0; i < 1 <<uint(h+w); i++ {
		cnt := 0
		for s := 0; s < h; s++ {
			for t := 0; t < w; t++ {
				fmt.Println(i&(1<<uint(s)))
				if i&(1<<uint(s)) > 0 {
						continue
				}
				if i&(1<<uint(h+t)) > 0 {
					continue
				}
				if js[s][t] == '#' {
					cnt++
				}
			}
		}
		if cnt == k {
			ans++
		}
	}
	fmt.Println(ans)
}
