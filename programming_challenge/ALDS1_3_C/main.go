package main

import (
	"container/list"
	"fmt"
)

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/all/ALDS1_3_B
// キュー

type queue struct {
	list *list.List
}

type proc struct {
	name string
	time int
}

func (q *queue) enqueue(p proc) {
	q.list.PushBack(p)
}

func (q *queue) dequeue() proc {
	return q.list.Remove(q.list.Front()).(proc)
}

func main() {
	var n, qt int
	fmt.Scan(&n, &qt)

	q := queue{list: list.New()}
	for i := 0; i < n; i++ {
		var name string
		var time int
		fmt.Scan(&name, &time)
		q.enqueue(proc{name: name, time: time})
	}

	var sum int
	for 0 < q.list.Len() {
		v := q.dequeue()
		time := min(v.time, qt)
		v.time -= time
		sum += time
		if 0 < v.time {
			q.enqueue(v)
		} else {
			fmt.Printf("%s %d\n", v.name, sum)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
