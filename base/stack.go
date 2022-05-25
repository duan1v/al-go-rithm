package base

import "fmt"

// 基于链表的栈
type Stack struct {
	size     int
	headNode *NodeSL
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Unshift(data interface{}) int {
	node := &NodeSL{Data: data}
	node.NextNode = s.headNode
	s.headNode = node
	s.size++
	return s.size
}

func (s *Stack) Shift() interface{} {
	x := s.headNode
	if x != nil {
		s.headNode = x.NextNode
		s.size--
	}
	return x.Data
}

func (s *Stack) Peek() interface{} {
	return s.headNode.Data
}

func TestStack() {
	s := Stack{}
	fmt.Println("s是空栈 ", s.IsEmpty())
	s.Unshift(1)
	s.Unshift(3)
	s.Unshift(2)
	s.Unshift(4)
	s.Unshift(6)
	x := s.Size()
	y := (s.Peek()).(int)
	z := s.Size()
	o := s.Shift().(int)
	p := s.Size()
	fmt.Printf(
		"s的长度是%d,peek一个数:%d,之后,s的长度是%d,shift一个数:%d,之后,s的长度是%d\n",
		x, y, z, o, p,
	)

	for !s.IsEmpty() {
		n, _ := (s.Shift()).(int)
		fmt.Printf("%v ", n)
	}
}

// ====================================================================================

// 基于切片的栈

type SStack[T interface{}] struct {
	elements []T
}

func (ss *SStack[T]) IsEmpty() bool {
	return len(ss.elements) == 0
}

func (ss *SStack[T]) Size() int {
	return len(ss.elements)
}

func (ss *SStack[T]) Unshift(t T) {
	ss.elements = append(ss.elements, t)
}

func (ss *SStack[T]) Shift() (t T, ok bool) {
	l := len(ss.elements)
	if l == 0 {
		return
	} else {
		t = ss.elements[l-1]
		ss.elements = ss.elements[0 : l-1]
		return t, true
	}
}

func (ss *SStack[T]) Peek() (t T, ok bool) {
	l := len(ss.elements)
	if l == 0 {
		return
	} else {
		t = ss.elements[l-1]
		return t, true
	}
}

func (ss *SStack[T]) Print() {
	fmt.Printf("%#v \n", ss.elements)
}

func TestSStack() {
	s := SStack[int]{}
	fmt.Println("s是空栈 ", s.IsEmpty())
	s.Unshift(1)
	s.Unshift(3)
	s.Unshift(2)
	s.Unshift(4)
	s.Unshift(6)
	x := s.Size()
	y, _ := s.Peek()
	z := s.Size()
	o, _ := s.Shift()
	p := s.Size()
	fmt.Printf(
		"s的长度是%d,peek一个数:%d,之后,s的长度是%d,shift一个数:%d,之后,s的长度是%d\n",
		x, y, z, o, p,
	)
	s.Print()
	for n, ok := s.Shift(); ok; n, ok = s.Shift() {
		fmt.Printf("%v \n", n)
	}
	s.Print()
}
