package main

import (
	"fmt"
	"math"
)

// コッホ曲線
// https://onlinejudge.u-aizu.ac.jp/problems/ALDS1_5_C

type point struct {
	x float64
	y float64
}
var th = math.Pi * 60.0 / 180.0

func koch(l point, r point, x int) {
	if x == 0 {
		fmt.Printf("%f %f\n", l.x, l.y)
	} else {
		p1 := point{
			x: (2.0*l.x + 1.0*r.x) / 3.0,
			y: (2.0*l.y + 1.0*r.y) / 3.0,
		}
		p2 := point{
			x: (1.0*l.x + 2.0*r.x) / 3.0,
			y: (1.0*l.y + 2.0*r.y) / 3.0,
		}
		m := point{
			x: (p2.x-p1.x)*math.Cos(th) - (p2.y-p1.y)*math.Sin(th) + p1.x,
			y: (p2.x-p1.x)*math.Sin(th) + (p2.y-p1.y)*math.Cos(th) + p1.y,
		}
		koch(l, p1, x-1)
		koch(p1, m, x-1)
		koch(m, p2, x-1)
		koch(p2, r, x-1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	a := point{x: 0, y:0}
	b := point{x: 100, y:0}

	fmt.Println(a.x, a.y)
	koch(a, b, n)
	fmt.Println(b.x, b.y)
}

