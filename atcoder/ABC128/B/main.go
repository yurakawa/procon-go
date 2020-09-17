package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// sort struct
func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	es := make(entries, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s := sc.Text()

		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		e := entry{s, p, i + 1}
		es[i] = e
	}
	sort.Sort(es)
	for i := 0; i < n; i++ {
		fmt.Println(es[i].id)
	}
}

//////
// for sort
type entry struct {
	key   string
	value int
	id    int
}

type entries []entry

func (es entries) Len() int {
	return len(es)
}

func (es entries) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}
func (es entries) Less(i, j int) bool {
	if es[i].key != es[j].key {
		return es[i].key < es[j].key
	}
	return es[i].value > es[j].value
}
