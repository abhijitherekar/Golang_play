package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func isValid(s string) bool {
	brack := stack.New()
	fmt.Println("******", s)
	//brack.Push(s[0])
	for i, _ := range s {
		switch s[i] {
		case ')':
			fmt.Println("taking : ", brack.Peek())
			if brack.Pop() != '(' {
				fmt.Println("return false ***", brack.Len(), brack.Peek(), brack.Pop())
				return false
			}
			continue
		case '}':
			fmt.Println("taking : ", brack.Peek())
			if brack.Pop() != '{' {
				return false
			}
			continue
		case ']':
			fmt.Println("taking :", brack.Peek())
			if brack.Pop() != '[' {
				return false
			}
			continue
		}
		fmt.Println("pushing: ", rune(s[i]))
		brack.Push(rune(s[i]))
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))

	fmt.Println(isValid("()[]{}"))

	fmt.Println(isValid("(]"))

	fmt.Println(isValid("([)]"))

	fmt.Println(isValid("{[]}"))
}
