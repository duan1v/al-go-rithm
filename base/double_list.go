package base

import (
	"fmt"
	"strings"
)

type NodeDL struct {
	Data     interface{}
	PrevNode *NodeDL
	NextNode *NodeDL
}

type DoubleList struct {
	size     int
	headNode *NodeDL
}

func (dl *DoubleList) Insert(index int, data interface{}) int {
	if index < 1 || index > dl.size+1 {
		panic("索引超过长度限制!")
	}
	node := NodeDL{Data: data}
	if dl.size == 0 {
		dl.headNode = &NodeDL{}
	}
	cn := dl.headNode
	if dl.size == 0 {
		node.PrevNode = cn
		cn.NextNode = &node
		goto END
	}
	for i := 1; i <= dl.size; i++ {
		nn := cn.NextNode
		if i == index {
			node.PrevNode = cn
			nn.PrevNode = &node
			node.NextNode = nn
			cn.NextNode = &node
			goto END
		}
		cn = nn
	}
	if index == dl.size+1 {
		node.PrevNode = cn
		cn.NextNode = &node
		goto END
	}
END:
	dl.size++
	return dl.size
}

func (dl *DoubleList) Print() {
	fmt.Printf("\n%s\n%s\n", strings.Repeat("=", 20), "开始打印该链表各个节点的数据:")
	node := dl.headNode
	for i := 1; i <= dl.size; i++ {
		node = node.NextNode
		fmt.Printf("%#v\n", node.Data)
	}
	fmt.Println(strings.Repeat("=", 20))
}

func (dl *DoubleList) validateIndex(index int) {
	if index < 1 || index > dl.size {
		panic("索引位置超过长度!")
	}
}

func (dl *DoubleList) PrintWithIndex(index int) {
	dl.validateIndex(index)
	node := dl.headNode
	for i := 1; i <= dl.size; i++ {
		node = node.NextNode
		if index == i {
			fmt.Printf("%#v\n", node.Data)
			break
		}
	}
}

func (dl *DoubleList) Delete(index int) int {
	dl.validateIndex(index)
	node := dl.headNode
	for i := 0; i < dl.size; i++ {
		nn := node.NextNode
		if index == i+1 {
			nn.NextNode.PrevNode = node
			node.NextNode = nn.NextNode
			dl.size--
			break
		}
		node = nn
	}
	return dl.size
}

func (dl *DoubleList) Size() int {
	return dl.size
}

func (dl *DoubleList) Reverse() {
	cn := dl.headNode.NextNode
	pn := &NodeDL{}
	for i := 1; i <= dl.size; i++ {
		nn := cn.NextNode
		cn.NextNode = pn
		pn = cn
		cn = nn
	}
	dl.headNode.NextNode = pn
	pn.PrevNode = dl.headNode
}

func TestDoubleList() {
	dl := DoubleList{}
	dl.Insert(1, 1)
	dl.Insert(2, 2)
	dl.Insert(3, 3)
	dl.PrintWithIndex(2)
	fmt.Println("dl的长度是: ", dl.Size())
	dl.Print()
	dl.Insert(2, 3)
	dl.Insert(1, 3)
	dl.Print()
	dl.PrintWithIndex(2)
	dl.Delete(3)
	// dl.Delete(2)
	dl.Print()
	dl.Reverse()
	dl.Print()
	fmt.Println("dl的长度是: ", dl.Size())
}
