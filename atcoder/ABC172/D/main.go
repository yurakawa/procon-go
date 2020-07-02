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

	a := make(map[int]int, 1e5+1)
	var sum uint64
	for i:= 0; i < n; i++ {
		tmp := sc.Int()
		a[tmp]++
		sum += uint64(tmp)
	}

	q := sc.Int()

	for i:= 0; i < q; i++ {
		b, c := sc.Int(), sc.Int()
		sum -= uint64(a[b]*b - a[b]*c)
		a[c] += a[b]
		a[b] = 0

		fmt.Println(sum)
	}
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

func lower_bound(arr []int, value int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < value {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
