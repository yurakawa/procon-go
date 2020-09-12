package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Union-Find
// https://atcoder.jp/contests/atc001/tasks/unionfind_a

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	uf := newUnionFind(n)
	for i := 0; i < q; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())

		if p == 0 {
			// 連結
			uf.unite(a, b)

		} else {
			// 判定

			if uf.same(a, b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}

	}

}

//////
// Union-Find
type UnionFind struct {
	par []int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

// xの所属するグループを返す
func (u UnionFind) root(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.root(u.par[x])
	return u.par[x]
}

// xの所属するグループ と yの所属するグループ をJoinする
func (u UnionFind) unite(x, y int) {
	xr := u.root(x)
	yr := u.root(y)
	if xr == yr {
		return
	}
	u.par[yr] += u.par[xr]
	u.par[xr] = yr
}

// xとyが同じグループに所属するかどうかを返す
func (u UnionFind) same(x, y int) bool {
	return u.root(x) == u.root(y)
}

// xの所属するグループの木の大きさを返す
func (u UnionFind) size(x int) int {
	return -u.par[u.root(x)]
}
