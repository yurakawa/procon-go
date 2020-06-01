package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/all/ALDS1_3_C
// 連結リスト


func main() {
	var n int
	fmt.Scan(&n)

	initNode()
	for i := 0; i < n; i++ {
		var c string
		var x int
		fmt.Scan(&c)

		switch c {
		case "insert":
			fmt.Scan(&x)
			insert(x)
		case "delete":
			fmt.Scan(&x)
			deleteKey(x)
		case "deleteFirst":
			deleteFirst()
		case "deleteLast":
			deleteLast()
		default:
			panic("err")
		}
	}
	fmt.Println(printList())

}

var dummy = &node{} // 番兵
type node struct {
	prev, next *node
	key int

}

func initNode() {
	dummy.next= dummy
	dummy.prev= dummy
}

// insert insets the new node at the top position of node
func insert(key int) {
	new := &node{key:key}
	new.next = dummy.next
	dummy.next.prev = new
	dummy.next= new
	new.prev = dummy
}

// search finds and returns the first node containing the specified key
func listSearch(key int) *node {
	current := dummy.next
	for  current.key != key {
		if current == dummy {
			return nil
		}
		current = current.next
	}
	return current
}

func deleteNode(n *node) {
	if n == nil {
		return // 番兵の場合処理しない
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n = nil
}

func deleteFirst() {
	deleteNode(dummy.next)
}

func deleteLast() {
	deleteNode(dummy.prev)
}

func deleteKey(key int) {
	deleteNode(listSearch(key))
}

func printList() string {
	buf := bytes.Buffer{}
	current := dummy.next
	for {
		buf.WriteString(strconv.Itoa(current.key))
		current = current.next
		if current == dummy {
			break
		}
		buf.WriteString(" ")
	}
	return buf.String()
}
