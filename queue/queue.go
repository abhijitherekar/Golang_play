package main

import (
	"fmt"
)

type (
	node struct {
		val  interface{}
		next *node
	}
	queue struct {
		front  *node
		last   *node
		length int
	}
)

func (q *queue) enqueue(val interface{}) {
	if q.length == 0 {
		fmt.Println("\n in enque")
		n := &node{val: val, next: nil}
		q.front = n
		q.last = n
		q.length++
		return
	}
	n := &node{val: val}
	q.last.next = n
	q.last = n
	q.length++
}

func (q *queue) leng() int {
	return q.length
}

func main() {
	q := &queue{front: nil, last: nil, length: 0}
	q.enqueue(14)
	q.enqueue(4)
	fmt.Println(q.leng())
}
