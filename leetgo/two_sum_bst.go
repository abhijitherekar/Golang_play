package main

import (
	"fmt"
)

type (
	Tree struct {
		root   *node
		size   int
		height int
		nchild int
	}
	node struct {
		left  *node
		right *node
		data  int
	}
)

func newnode(data int) *node {
	root := new(node)
	root.data = data
	root.left = nil
	root.right = nil
	return root
}

/*
func inorder_append(root *node, a []int) []int {

	if root == nil {
		return nil
	} else {
		//a = append(a, inorder_append(root.left, a), inorder_append(root.right, a))
		a = inorder_append(root.left, a)
		a = append(a, root.data)
		a = append(a, inorder_append(root.right, a)...)
		return a
	}
}
*/
func (n *node) insert(data int) *node {
	if n == nil {
		return newnode(data)
	}
	if data < n.data {
		if n.left == nil {
			n.left = n.left.insert(data)
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = n.right.insert(data)
		} else {
			n.right.insert(data)
		}
	}
	return nil
}

func (t *Tree) insert_bst(data int) {
	if t.root == nil {
		t.root = newnode(data)
		t.size++
		return
	}
	t.size++
	t.root.insert(data)
}

func print_inorder(root *node) {
	if root == nil {
		//fmt.Println("root is nil")
		return
	}

	if root != nil {
		print_inorder(root.left)
		fmt.Println("**", root.data)
		print_inorder(root.right)
	}
}

func (t *Tree) Size() int {
	return t.size
}

func sizeroot(root *node) int {
	if root == nil {
		return 0
	}
	return sizeroot(root.left) + 1 + sizeroot(root.right)
}

func New() *Tree {
	return &Tree{}
}

func maxdepth(root *node) int {
	if root == nil {
		return 0
	}
	rdepth := maxdepth(root.right)
	ldepth := maxdepth(root.left)

	if rdepth > ldepth {
		return 1 + rdepth
	} else {
		return 1 + ldepth
	}
}

func (t *Tree) minVal() int {
	if t.root == nil {
		return -1
	}
	node := t.root
	for node.left != nil {
		node = node.left
	}
	return node.data
}

func hasPathSum(root *node, sum int) bool {
	if root == nil || sum == 0 {
		return false
	}

	if sum-root.data == 0 {
		return true
	}
	return hasPathSum(root.left, sum-root.data) || hasPathSum(root.right, sum-root.data)
}
func printpath(a []int, root *node) {
	if root == nil {
		return
	}
	fmt.Println("\n a before append n print:", a, "\n")
	if root.left == nil && root.right == nil {
		a = append(a, root.data)
		fmt.Println(a)
		return
	}
	fmt.Println("appending: ", root.data)
	a = append(a, root.data)
	printpath(a, root.left)
	if root.right != nil {
		fmt.Println("\n now going to right with", a, "towards: ", root.right.data)
		printpath(a, root.right)
	}
}
func printpaths(root *node) {
	a := make([]int, 10, 1000)
	//a = append(a, root.data)
	printpath(a, root)
}

func vertsum(root *node, sum int) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		sum += root.data
		fmt.Println(sum)
		return
	}
	sum += root.data
	vertsum(root.left, sum)
	vertsum(root.right, sum)
}

func depth(root *node, val int, level int) int {
	if root == nil {
		return 0
	}
	if root.data == val {
		return level
	}
	if a := depth(root.left, val, level+1); a != 0 {
		return a
	}
	return depth(root.right, val, level+1)
}

func (t *Tree) leveltraversal(root *node) []int {
	lvl := make([]int, 10)
	//for i := range lvl {
	//	lvl[i] = make([]int, 0, 10)
	//}

	if root == nil {
		return lvl
	}
	q := &queue{front: nil, end: nil, length: 0}
	q.enqueue(root)
	for q.size() != 0 {
		temp := q.dqueue().(*node)
		idx := depth(t.root, temp.data, 0)
		fmt.Println(idx)
		lvl[idx] = (lvl[idx] + temp.data)
		if temp.left != nil {
			q.enqueue(temp.left)
		}
		if temp.right != nil {
			q.enqueue(temp.right)
		}
	}
	return lvl
}

func (t *Tree) zigzagtravel(root *node) [][]int {
	lvl := make([][]int, 10)
	for i := range lvl {
		lvl[i] = make([]int, 0, 10)
	}

	if root == nil {
		return lvl
	}
	s1 := &stack{top: nil, length: 0}
	s2 := &stack{top: nil, length: 0}
	s1.push(root)
	for s1.size() != 0 || s2.size() != 0 {
		for s1.size() != 0 {
			temp := s1.pop().(*node)
			idx := depth(t.root, temp.data, 0)
			lvl[idx] = append(lvl[idx], temp.data)
			if temp.left != nil {
				s2.push(temp.left)
			}
			if temp.right != nil {
				s2.push(temp.right)
			}
		}
		for s2.size() != 0 {
			temp := s2.pop().(*node)
			idx := depth(t.root, temp.data, 0)
			lvl[idx] = append(lvl[idx], temp.data)
			if temp.right != nil {
				s1.push(temp.right)
			}
			if temp.left != nil {
				s1.push(temp.left)
			}
		}
	}
	return lvl
}

type (
	item struct {
		val  interface{}
		prev *item
	}
	stack struct {
		top    *item
		length int
	}
)

func (s *stack) size() int {
	return s.length
}

func (s *stack) push(val interface{}) {
	if s.length == 0 {
		n := &item{val: val, prev: nil}
		s.top = n
		s.length++
	} else {
		n := &item{val: val, prev: s.top}
		s.top = n
		s.length++
	}
}

func (s *stack) pop() interface{} {
	if s.length == 0 {
		return nil
	} else {
		temp := s.top
		s.top = temp.prev
		s.length--
		return temp.val
	}
}

type (
	elem struct {
		val  interface{}
		next *elem
	}
	queue struct {
		front  *elem
		end    *elem
		length int
	}
)

func (q *queue) size() int {
	return q.length
}
func (q *queue) enqueue(val interface{}) {
	if q.length == 0 {
		n := &elem{val: val, next: nil}
		q.front = n
		q.end = n
		q.length++
	} else {
		n := &elem{val: val, next: nil}
		q.end.next = n
		q.end = n
		q.length++
	}
}

func (q *queue) dqueue() interface{} {
	if q.length == 0 {
		return nil
	} else {
		temp := q.front
		q.front = temp.next
		q.length--
		return temp.val
	}
}

func main() {
	t := New()
	//a := make([]int, 10, 1000)
	t.insert_bst(50)
	t.insert_bst(60)
	t.insert_bst(40)
	t.insert_bst(20)

	t.insert_bst(30)
	t.insert_bst(10)
	t.insert_bst(55)
	t.insert_bst(70)

	t.insert_bst(45)
	t.insert_bst(80)
	fmt.Println(t.Size(), sizeroot(t.root), maxdepth(t.root), t.minVal(), hasPathSum(t.root, 22), hasPathSum(t.root, 11))
	//insert_bst(root, 10)
	printpaths(t.root)
	print_inorder(t.root)
	//a = inorder_append(root, a)
	fmt.Println("level sum :", t.leveltraversal(t.root))
	fmt.Println(t.zigzagtravel(t.root))
	//fmt.Println(a)
}
