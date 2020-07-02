package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
func main() {
	sc := NewBufScanner(os.Stdin)
	n, m, k := sc.Int(), sc.Int(), sc.Int()


	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		tmp := sc.Int()
		a[i] = a[i-1] + tmp
	}

	b := make([]int ,m+1)
	for i := 1; i <= m; i++ {
		tmp := sc.Int()
		b[i] = b[i-1] + tmp
	}



	ans := 0

	j := m
	for i :=0; i <=n; i++ {
		if a[i] > k {
			continue
		}
		for b[j] > k - a[i] {
			j--
		}
		ans = max(ans, i + j)
	}

	fmt.Println(ans)
}


type BufScanner struct {
	*bufio.Scanner
}

func NewBufScanner(r io.Reader) *BufScanner {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	// s.Split(scanWords)
	return &BufScanner{s}
}

func scanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for isSpace(data[start]) {
		start++
	}
	for i := start; i < len(data); i++ {
		if isSpace(data[i]) {
			return i+1, data[start:i], nil
		}
	}
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}
func isSpace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\r' || b == '\t'
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
