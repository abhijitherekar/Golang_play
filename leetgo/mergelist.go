package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type (
	List struct {
		head   *ListNode
		length int
	}
	ListNode struct {
		Val  int
		Next *ListNode
	}
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *List {
	ml := New()
	for {
		//fmt.Println("hhhhh")
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil && l2 != nil {
			ml.Insert(l2.Val)
			l2 = l2.Next
			continue
		}

		if l2 == nil && l1 != nil {
			ml.Insert(l1.Val)
			l1 = l1.Next
			continue
		}
		if l1.Val <= l2.Val {
			ml.Insert(l1.Val)
			l1 = l1.Next
		} else if l2.Val < l1.Val {
			ml.Insert(l2.Val)
			l2 = l2.Next
		}
	}
	//ml.Print()
	return ml
}

func New() *List {
	return &List{}
}

func (l *List) Insert(vals ...interface{}) {
	for _, val := range vals {
		if l.length == 0 {
			head := &ListNode{Val: val.(int)}
			head.Next = nil
			l.head = head
			l.length++
		} else {
			curr := l.head
			for curr.Next != nil {
				curr = curr.Next
			}
			curr.Next = &ListNode{Val: val.(int)}
			l.length++
		}
	}
}

func (l *List) Print() {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func main() {

	l1 := New()
	l1.Insert(1, 2, 4)
	l2 := New()
	l2.Insert(3, 5, 6)
	//l1.Print()
	//l2.Print()
	mergeTwoLists(l1.head, l2.head).Print()

}
