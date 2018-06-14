package stack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		val  interface{}
		prev *node
	}
)

func New() *Stack {
	return &Stack{}
}

func (this *Stack) Push(val interface{}) {
	if this.length == 0 {
		this.top = &node{val, nil}
		this.length++
	} else {
		n := &node{val, this.top}
		this.top = n
		this.length++
	}
}

func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	val := this.top.val
	this.top = this.top.prev
	this.length--
	return val
}

func (this *Stack) Peek() interface{} {
	return this.top.val
}

func (this *Stack) Length() int {
	return this.length
}
