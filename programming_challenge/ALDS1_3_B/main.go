package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	s [1000]int
	top int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var stack Stack
	sc.Scan()
	ss := sc.Text()
	for _, s := range strings.Split(ss, " ") {
		switch s {
		case "*" :
			r := stack.pop()
			l := stack.pop()
			stack.push(l*r)
		case "+" :
			r := stack.pop()
			l := stack.pop()
			stack.push(l+r)
		case "-" :
			r := stack.pop()
			l := stack.pop()
			stack.push(l-r)
		default:
			num,_ := strconv.Atoi(s)
			stack.push(num)
		}
	}
	fmt.Println(stack.pop())
}


func (stack *Stack)push(x int) {
	stack.top++
	stack.s[stack.top] = x
}

func (stack *Stack)pop() int{
	x := stack.s[stack.top]
	stack.top--
	return x
}
