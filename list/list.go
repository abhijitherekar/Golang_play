package list

type List struct {
	head   *Element
	length int
}

type Element struct {
	next, prev *Element
	Val        interface{}
}

func (l List) init() {
	l.head.next = l.head
	l.head.prev = l.head
}

func New() *List {
	var l List
	l.head = new(Element)
	l.length = 0
	l.init()
	return &l
}

func (l *List) Insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	l.length++
	return e
}

//push the element to the front of the list
func (l *List) PushToFront(val interface{}) *Element {
	e := &Element{prev: nil, next: nil, Val: val}
	return l.Insert(e, l.head)
}

func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
	l.length--
	return e
}

func (l *List) Remove(e *Element) interface{} {
	if e != nil {
		e = l.remove(e)
		return e.Val
	}
	return nil
}

func (l *List) MoveToFront(e *Element) *Element {
	return l.PushToFront(l.Remove(e))
}

func (l *List) Size() int {
	return l.length
}

func (l *List) Front() *Element {
	if l.length == 0 {
		return nil
	}
	return l.head.next
}

func (l *List) Back() *Element {
	if l.length != 0 {
		return l.head.prev
	}
	return nil
}

//return the previous element
func (e *Element) Prev() *Element {
	if e != nil && e.prev != nil {
		return e.prev
	}
	return nil
}

//return the next element
func (e *Element) Next() *Element {
	if e != nil && e.next != nil {
		return e.next
	}
	return nil
}
