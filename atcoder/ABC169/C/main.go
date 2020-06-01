package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a int
	var b string
	fmt.Scan(&a, &b)
	bn, _ := strconv.Atoi(strings.Replace(b,".", "", -1))

	fmt.Println(a*bn/100)

}

