package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	sc := NewBufScanner(os.Stdin)
	n := sc.Int()

	dic := make(map[string]int, n)
	for i := 0; i < n; i++ {
		command, address := sc.String(), sc.String()
		if command == "insert" {
			dic[address] = 1
		}

		if command == "find" {
			if dic[address] == 1 {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}
	}
}

type BufScanner struct {
	*bufio.Scanner
}

func NewBufScanner(r io.Reader) *BufScanner {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	return &BufScanner{s}
}

func (s *BufScanner) String() string {
	s.Scan()
	return s.Text()
}

func (s *BufScanner) Int() int {
	return int(s.UInt())
}

func (s *BufScanner) UInt() uint64 {
	s.Scan()
	i, e := strconv.ParseUint(s.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

