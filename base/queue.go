package base

import (
	"fmt"
)

// 基于链表的队列
type Queue struct {
	size int
	head *NodeSL
	tail *NodeSL
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Push(data interface{}) int {
	node := NodeSL{Data: data}
	if q.size == 0 {
		q.head = &node
		q.tail = &node
	} else {
		q.tail.NextNode = &node
		q.tail = &node
	}
	q.size++
	return q.size
}

func (q *Queue) Pop() (ans *NodeSL) {
	if q.size == 0 {
		panic("当前队列为空!")
	} else {
		ans = q.head
		q.head = q.head.NextNode
		q.size--
	}
	if q.head == nil {
		q.tail = nil
	}
	return
}

func (q *Queue) Peek() (ans *NodeSL) {
	if q.head != nil {
		ans = q.head
	}
	return
}

func TestQueue() {
	q := Queue{}
	fmt.Println("q是空队列 ", q.IsEmpty())
	q.Push(1)
	q.Push(3)
	q.Push(2)
	y := q.Size()
	x := q.Peek()
	z := q.Size()
	o := q.Pop()
	p := q.Size()
	fmt.Printf(
		"q的长度是%d,peek一个数:%d,之后,q的长度是%d,pop一个数:%d,之后,q的长度是%d\n",
		y, x.Data, z, o.Data, p,
	)
}

// ==============================================================================

// 基于切片的对列

type SQueue[T interface{}] struct {
	elements []T
}

func (sq *SQueue[T]) IsEmpty() bool {
	return len(sq.elements) == 0
}

func (sq *SQueue[T]) Push(t T) {
	sq.elements = append(sq.elements, t)
}

func (sq *SQueue[T]) Pop() (T, bool) {
	var t T
	if len(sq.elements) == 0 {
		return t, false
	}
	t = sq.elements[0]
	sq.elements = sq.elements[1:]
	return t, true
}

func (sq *SQueue[T]) Peek() (t T, ok bool) {
	if len(sq.elements) == 0 {
		return t, false
	}
	return sq.elements[0], true
}

func (sq *SQueue[T]) Size() int {
	return len(sq.elements)
}

func (sq *SQueue[T]) Print() {
	fmt.Printf("%#v \n", sq.elements)
}

func TestSQueue() {
	q := SQueue[float32]{}
	fmt.Println("q是空队列 ", q.IsEmpty())
	q.Push(1)
	q.Push(3)
	q.Push(2)
	y := q.Size()
	x, _ := q.Peek()
	z := q.Size()
	o, _ := q.Pop()
	p := q.Size()
	// fmt.Printf(
	// 	"q的长度是%d,peek一个数:%d,之后,q的长度是%d,pop一个数:%d,之后,q的长度是%d\n",
	// 	y, x, z, o, p,
	// )
	fmt.Printf(
		"q的长度是%d,peek一个数:%f,之后,q的长度是%d,pop一个数:%f,之后,q的长度是%d\n",
		y, x, z, o, p,
	)
	q.Print()
	for n, ok := q.Pop(); ok; n, ok = q.Pop() {
		fmt.Printf("%v \n", n)
	}
	q.Print()
}
