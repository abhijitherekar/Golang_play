package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := New()
	fmt.Println(l)
	// Single element list
	e := l.PushToFront(1)
	e1 := l.PushToFront(2)
	e2 := l.PushToFront(3)
	e3 := l.PushToFront(8)
	fmt.Println(l, e, e1, e2, e3)
	fmt.Println(l.Front().val.(int)) //8
	l.MoveToFront(e)
	fmt.Println(l) //1
	e1 = l.MoveToFront(e1)
	fmt.Println(l.Front().val.(int))
	l.Remove(e1)
	fmt.Println(l.Front().val.(int))
}
