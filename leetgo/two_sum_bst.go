package main

import (
	"fmt"
)

type node struct {
	left  *node
	right *node
	data  int
}

func newnode(data int) *node {
	root := new(node)
	root.data = data
	root.left = nil
	root.right = nil
	return root
}

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

func insert_bst(root *node, data int) *node {
	if root == nil {
		return newnode(data)
	}
	if data < root.data {
		root.left = insert_bst(root.left, data)
	} else {
		root.right = insert_bst(root.right, data)
	}
	return root
}

func print_inorder(root *node) {
	if root == nil {
		return
	}

	if root != nil {
		print_inorder(root.left)
		fmt.Println(root.data)
		print_inorder(root.right)
	}
}

func main() {
	var root *node
	a := make([]int, 10, 1000)
	root = insert_bst(root, 5)
	insert_bst(root, 6)
	insert_bst(root, 4)
	insert_bst(root, 2)
	//insert_bst(root, 10)
	print_inorder(root)
	a = inorder_append(root, a)
	fmt.Println(a)
}
