package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	w := []string{"Sunny", "Cloudy", "Rainy"}
	for k, v := range w{
		if s == v  {
			fmt.Println(w[(k+1)%3])
			return
		}
	}
}
