package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

// B - Popular Vote
// https://atcoder.jp/contests/abc161/tasks/abc161_b

func main() {
	scanner := NewBufScanner(os.Stdin)
	N, M := scanner.Int(), scanner.Float64()
	var con float64
	var as []float64
	var sum float64
	for i := 0; i < N; i++ {
		a := scanner.Float64()
		as = append(as, a)
		sum += a
	}

	for _, a := range as {
		if sum > a * 4 * M {
			continue

		}
		con++
	}
	if con  >= M {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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

func (s *BufScanner) Float64() float64 {
	s.Scan()
	i, e := strconv.ParseFloat(s.Text(),64)
	if e != nil {
		panic(e)
	}
	return i
}

func (s *BufScanner) ArrayInt(n int) (r []int) {
	for i := 0; i < n; i++ {
		r = append(r, s.Int())
	}
	return
}

///////////////


// FullScanner for stdin
type FullScanner struct {
	*bytes.Buffer
}

func NewFullScanner(r io.Reader) *FullScanner {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(b)
	return &FullScanner{buf}
}

func (s *FullScanner) Int() (r int) {
	fmt.Fscan(s, &r)
	return
}

// Scanner for stdin
type Scanner struct {
}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Text() (r string) {
	fmt.Scan(&r)
	return
}

func (s *Scanner) UInt() (r uint64) {
	fmt.Scan(&r)
	return
}

func (s *Scanner) Int() (r int) {
	fmt.Scan(&r)
	return
}

func (s *Scanner) UInt32() (r uint32) {
	fmt.Scan(&r)
	return
}

func (s *Scanner) Int32() (r int32) {
	fmt.Scan(&r)
	return
}

func (s *Scanner) Double() (r float64) {
	fmt.Scan(&r)
	return
}

func (s *Scanner) ArrayText(n int) (r []string) {
	for i := 0; i < n; i++ {
		r = append(r, s.Text())
	}
	return
}

func (s *Scanner) ArrayInt(n int) (r []int) {
	for i := 0; i < n; i++ {
		r = append(r, s.Int())
	}
	return
}

func (s *Scanner) ArrayDouble(n int) (r []float64) {
	for i := 0; i < n; i++ {
		r = append(r, s.Double())
	}
	return
}

func arrayIntToArrayString(arr []int) (r []string) {
	for _, v := range arr {
		r = append(r, fmt.Sprintf("%d", v))
	}
	return
}

func arrayStringToArrayInt(arr []string) (r []int) {
	for _, v := range arr {
		i, e := strconv.ParseInt(v, 10, 64)
		if e != nil {
			panic(e)
		}
		r = append(r, int(i))
	}
	return
}

func Sum(arr ...int) (sum int) {
	for _, x := range arr {
		sum += x
	}
	return
}


// https://gobyexample.com/collection-functions
func AllInt(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

