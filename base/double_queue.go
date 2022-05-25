package base

import "fmt"

type DoubleQueue struct {
	size int
	head *NodeDL
	tail *NodeDL
}

func (dq *DoubleQueue) IsEmpty() bool {
	return dq.size == 0
}

func (dq *DoubleQueue) Size() int {
	return dq.size
}

// 进来的节点放在尾巴上
func (dq *DoubleQueue) Push(data interface{}) int {
	node := &NodeDL{Data: data, PrevNode: dq.tail}
	if dq.head == nil {
		dq.head = node
		dq.tail = node
	} else {
		dq.tail.NextNode = node
		dq.tail = node
	}
	dq.size++
	return dq.size
}

// 进来的节点放头部
func (dq *DoubleQueue) Unshift(data interface{}) int {
	node := &NodeDL{Data: data, NextNode: dq.head}
	if dq.head == nil {
		dq.head = node
		dq.tail = node
	} else {
		dq.head.PrevNode = node
		dq.head = node
	}
	dq.size++
	return dq.size
}

// 从头部取出
func (dq *DoubleQueue) Shift() *NodeDL {
	hn := dq.head
	if hn != nil {
		hn.NextNode.PrevNode = nil
		dq.head = hn.NextNode
		dq.size--
	}
	return hn
}

// 从尾部取出
func (dq *DoubleQueue) Pop() *NodeDL {
	tn := dq.tail
	if tn != nil {
		tn.PrevNode.NextNode = nil
		dq.tail = tn.PrevNode
		dq.size--
	}
	return tn
}

func (dq *DoubleQueue) PeekHead() *NodeDL {
	return dq.head
}

func (dq *DoubleQueue) PeekTail() *NodeDL {
	return dq.tail
}

func TestDoubleQueue() {
	dq := DoubleQueue{}
	fmt.Println("dq是空的双向队列 ", dq.IsEmpty())
	dq.Push(1)
	dq.Push(2)
	dq.Push(3)
	dq.Unshift(-1)
	dq.Unshift(-2)
	dq.Unshift(-3)
	fmt.Println("dq的长度是 ", dq.Size())
	x := dq.Pop()
	fmt.Println("dq pop 了一个数 ", x.Data, "此时的长度是 ", dq.Size())
	y := dq.PeekTail()
	fmt.Println("dq peekTail 了一个数 ", y.Data, "此时的长度是 ", dq.Size())
	j := dq.Shift()
	fmt.Println("dq shift 了一个数 ", j.Data, "此时的长度是 ", dq.Size())
	k := dq.PeekHead()
	fmt.Println("dq peekHead 了一个数 ", k.Data, "此时的长度是 ", dq.Size())

}
