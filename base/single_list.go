package base

import (
	"fmt"
	"strings"
)

type NodeSL struct {
	Data     interface{}
	NextNode *NodeSL
}

func (ns *NodeSL) PrintWithHead() {
	fmt.Printf("\n%s\n%s\n", strings.Repeat("=", 20), "开始打印该节点及其之后节点的数据:")
	for ns != nil {
		fmt.Println(ns.Data)
		ns = ns.NextNode
	}
	fmt.Println(strings.Repeat("=", 20))
}

type SingleList struct {
	size     int
	headNode *NodeSL
}

func (sl *SingleList) Head() *NodeSL {
	if sl.headNode == nil {
		return nil
	}
	return sl.headNode.NextNode
}

func (sl *SingleList) Insert(index int, data interface{}) int {
	if index < 1 || index > sl.size+1 {
		return 0
	}
	node := NodeSL{Data: data}

	if sl.size == 0 {
		sl.headNode = &NodeSL{}
	}
	cn := sl.headNode
	for i := 1; i <= sl.size+1; i++ {
		if index == i {
			node.NextNode = cn.NextNode
			cn.NextNode = &node
			break
		}
		cn = cn.NextNode
	}

	sl.size++
	return sl.size
}

func (sl *SingleList) Print() {
	fmt.Printf("\n%s\n%s\n", strings.Repeat("=", 20), "开始打印该链表各个节点的数据:")
	node := sl.headNode.NextNode
	for i := 1; i <= sl.size; i++ {
		fmt.Printf("%#v\n", node.Data)
		node = node.NextNode
	}
	fmt.Println(strings.Repeat("=", 20))
}

func (sl *SingleList) Size() int {
	return sl.size
}

func (sl *SingleList) validateIndex(index int) {
	if !(index > 0 && index <= sl.size) {
		panic("链表长度不合法!")
	}
}

func (sl *SingleList) PrintWithIndex(index int) {
	sl.validateIndex(index)
	node := sl.headNode
	for i := 1; i <= sl.size; i++ {
		node = node.NextNode
		if index == i {
			fmt.Printf("%#v\n", node.Data)
			break
		}
	}
}

func (sl *SingleList) FindWithIndex(index int) *NodeSL {
	sl.validateIndex(index)
	node := sl.headNode
	for i := 1; i <= sl.size; i++ {
		node = node.NextNode
		if index == i {
			return node
		}
	}
	return nil
}

func (sl *SingleList) Update(index int, data int) {
	sl.validateIndex(index)
	node := sl.headNode
	for i := 1; i <= sl.size; i++ {
		node = node.NextNode
		if index == i {
			node.Data = data
			break
		}
	}
}

func (sl *SingleList) Delete(index int) int {
	sl.validateIndex(index)
	node := sl.headNode
	for i := 0; i < sl.size; i++ {
		if index == i+1 {
			node.NextNode = node.NextNode.NextNode
			sl.size--
			break
		}
		node = node.NextNode
	}
	return sl.size
}

func (sl *SingleList) Reverse() {
	if sl.size == 0 {
		return
	}
	// 最前面的空节点
	pn := &NodeSL{}
	cn := sl.headNode
	for i := 0; i <= sl.size; i++ {
		nn := cn.NextNode
		cn.NextNode = pn
		pn = cn
		cn = nn
	}
	sl.headNode.NextNode = pn
}

func TestSingleList() {
	sl := SingleList{}
	sl.Insert(1, 1)
	sl.Insert(2, 2)
	sl.Insert(3, 3)
	sl.Print()
	sl.PrintWithIndex(3)
	sl.Update(3, 8)
	sl.PrintWithIndex(3)
	sl.Insert(4, 4)
	sl.Delete(3)
	fmt.Print(sl.Size())
	sl.Print()
	sl.Reverse()
	sl.Print()
}
