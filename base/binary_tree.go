package base

import (
	"fmt"
	"unsafe"
)

type NodeBT struct {
	Data      interface{}
	LeftNode  *NodeBT
	RightNode *NodeBT
}

func (nb *NodeBT) Clear() {
	if nb != nil {
		nb.LeftNode.Clear()
		nb.LeftNode.Clear()
		nb = nil
	}
}

func (nb *NodeBT) Hight() int {
	if nb == nil {
		return 0
	} else {
		l := nb.LeftNode.Hight()
		r := nb.RightNode.Hight()
		if l > r {
			// 递归到这一层有节点,那就把这一层的高度1加进去
			return l + 1
		} else {
			return r + 1
		}
	}
}

func (nb *NodeBT) Size() int {
	if nb == nil {
		return 0
	} else {
		return nb.LeftNode.Size() + nb.RightNode.Size() + 1
	}
}

// 遍历pn及其子节点,与n进行比较
func (pn *NodeBT) FindParent(n *NodeBT) interface{} {
	if pn == nil {
		return nil
	}
	if uintptr(unsafe.Pointer(pn.LeftNode)) == uintptr(unsafe.Pointer(n)) ||
		uintptr(unsafe.Pointer(pn.RightNode)) == uintptr(unsafe.Pointer(n)) {
		return pn
	}
	l := pn.LeftNode.FindParent(n)
	if l != nil {
		return l
	} else {
		return pn.LeftNode.FindParent(n)
	}
}

// 深度遍历有前序、中序以及后序三种遍历方法
// 前序遍历
func (nb *NodeBT) PrePrint() {
	if nb != nil {
		fmt.Printf("%v ", nb.Data)
		nb.LeftNode.PrePrint()
		nb.RightNode.PrePrint()
	}
}

func (nb *NodeBT) PrePrintWithStack() {
	// s := SStack[*NodeBT]{}
	s := Stack{}
	for (!s.IsEmpty()) || nb != nil {
		if nb != nil {
			fmt.Printf("%v ", nb.Data)
			s.Unshift(nb)
			nb = nb.LeftNode
		} else {
			// sn, _ := s.Shift()
			// 注意断言类型及rn的类型;否则下面判断中的右节点读取不到,因为断言失败
			sn, ok := (s.Shift()).(*NodeBT)
			if !ok {
				panic("出栈数据不是二叉树节点!")
			}
			nb = sn.RightNode
		}
	}
}

// 中序遍历
func (nb *NodeBT) InPrint() {
	if nb != nil {
		nb.LeftNode.InPrint()
		fmt.Printf("%v ", nb.Data)
		nb.RightNode.InPrint()
	}
}

func (nb *NodeBT) InPrintWithStack() {
	s := SStack[*NodeBT]{}
	for (!s.IsEmpty()) || nb != nil {
		if nb == nil {
			sn, _ := s.Shift()
			fmt.Printf("%v ", sn.Data)
			nb = sn.RightNode
		} else {
			s.Unshift(nb)
			nb = nb.LeftNode
		}
	}
}

// 后序遍历
func (nb *NodeBT) PostPrint() {
	if nb != nil {
		nb.LeftNode.PostPrint()
		nb.RightNode.PostPrint()
		fmt.Printf("%v ", nb.Data)
	}
}

// 双栈法后序非递归
func (nb *NodeBT) PostPrintWithStack() {
	s := SStack[*NodeBT]{}
	q := SStack[*NodeBT]{}
	rn := nb
	for (rn != nil) || (!s.IsEmpty()) {
		if rn != nil {
			s.Unshift(rn)
			q.Unshift(rn)
			rn = rn.RightNode
		} else {
			sn, _ := s.Shift()
			rn = sn.LeftNode
		}
	}
	for !q.IsEmpty() {
		sn, _ := q.Shift()
		fmt.Printf("%#v ", sn.Data)
	}

}

// 这是一个顺着思维的后序非递归
func (nb *NodeBT) PostPrintWithStack1() {
	s := SStack[*NodeBT]{}
	rn := nb

	// 6.在1.中,如果栈中只有根节点时,
	// 1)如果满足2.或者3.,则栈s直接置空,又因rn==nil才能进到2./3.且rn没有再被赋值;所以直接跳出循环
	// 2)如果根节点的右子节点有子节点;则会在4.中出现死循环
	// 3)解决方法:手动置空栈;显然,根节点唯一出现在栈中的次数是可以确定的
	//   1>第一次是在压根节点入栈
	//   2>在有左子节点时,左子节点出栈后
	//   3>在2)的条件下,移除根节点的右节点之后
	//   4>由上面三个条件,可知,进栈一次;有左一次;有右时,于2),需要记一次(避免少打印或者死循环),于1),无所谓.
	c := 1
	if nb.LeftNode != nil {
		c++
	}
	if nb.RightNode != nil {
		c++
	}
	for (rn != nil) || (!s.IsEmpty()) {
		if rn != nil {
			s.Unshift(rn)
			rn = rn.LeftNode
		} else {
			// 这是一个查找不弹出的方法
			sn, _ := s.Peek()
			snr := sn.RightNode
			// 1.当前节点sn只有在:
			// 2.右子节点没有子节点时,先移除右子节点snr,再移除sn;
			if snr != nil && snr.LeftNode == nil && snr.RightNode == nil {
				fmt.Printf("%#v %#v ", snr.Data, sn.Data)
				s.Shift()
			} else if snr == nil {
				// 3.没有右子节点时,直接移除sn
				fmt.Printf("%#v ", sn.Data)
				s.Shift()
			} else {
				// 4.否则sn不可以被移除,继续将当前右子节点压栈
				rn = snr
			}
			// 5.依据先进后出,统计栈中仅有一个结点,即根结点的次数;
			if s.Size() == 1 {
				c--
			}
			if c == 0 {
				fmt.Printf("%#v ", nb.Data)
				break
			}
		}
	}
}

// 广度度优先遍历
func (nb *NodeBT) BreadthFirstPrint() {
	s := SStack[*NodeBT]{}
	s.Unshift(nb)
	for !s.IsEmpty() {
		sn, _ := s.Shift()
		fmt.Printf("%v ", sn.Data)
		// 先压右节点,则与先序遍历一致
		if sn.RightNode != nil {
			s.Unshift(sn.RightNode)
		}
		if sn.LeftNode != nil {
			s.Unshift(sn.LeftNode)
		}
	}
}

// 层次遍历
func (nb *NodeBT) LevelPrint() {
	q := SQueue[*NodeBT]{}
	q.Push(nb)
	for !q.IsEmpty() {
		n, _ := q.Pop()
		fmt.Printf("%#v ", n.Data)
		// 左节点先进,则顺序与生成树的数组一致
		if n.LeftNode != nil {
			q.Push(n.LeftNode)
		}
		if n.RightNode != nil {
			q.Push(n.RightNode)
		}
	}
}

// 这是一个完全二叉树
func (n *NodeBT) CreateOrResetWithArr(arr []int, index int) *NodeBT {
	if index < len(arr) {
		n = &NodeBT{Data: arr[index]}
		// 这种会用 &NodeBT{} 填上最后的节点
		// n.LeftNode = n.CreateOrResetWithArr(arr, 2*index+1)
		// n.RightNode = n.CreateOrResetWithArr(arr, 2*index+2)
		ln := n.CreateOrResetWithArr(arr, 2*index+1)
		if ln.Data != nil {
			n.LeftNode = ln
		}
		rn := n.CreateOrResetWithArr(arr, 2*index+2)
		if rn.Data != nil {
			n.RightNode = rn
		}
		return n
	}
	return &NodeBT{}
}

type BinaryTree struct {
	// Hight int // 无法在添加子树时获取,或者说每次添加时获取高度的开销较大;不如在需要时调用获取高度的函数
	Root *NodeBT
}

func (bt *BinaryTree) IsEmpty() bool {
	return bt.Root == nil
}

func (bt *BinaryTree) Clear(nb *NodeBT) {
	bt.Root.Clear()
}

func (bt *BinaryTree) Hight() int {
	return bt.Root.Hight()
}

func (bt *BinaryTree) Size() int {
	return bt.Root.Size()
}

func (bt *BinaryTree) FindParent(n *NodeBT) interface{} {
	if n == nil || uintptr(unsafe.Pointer(bt.Root)) == uintptr(unsafe.Pointer(n)) {
		return nil
	}
	return bt.Root.FindParent(n)
}

func (bt *BinaryTree) CreateWithArr(arr []int, index int) {
	bt.Root = (&NodeBT{}).CreateOrResetWithArr(arr, 0)
}

func TestBinaryTree() {
	//顺序存储的满二叉树或者完全二叉树
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bt := BinaryTree{}
	bt.CreateWithArr(arr, 0)
	fmt.Println("递归前序遍历：")
	bt.Root.PrePrint()
	fmt.Println()
	bt.Root.PrePrintWithStack()
	fmt.Println()
	fmt.Println("递归中序遍历: ")
	bt.Root.InPrint()
	fmt.Println()
	bt.Root.InPrintWithStack()
	fmt.Println()
	fmt.Println("递归后序遍历: ")
	bt.Root.PostPrint()
	fmt.Println()
	bt.Root.PostPrintWithStack()
	fmt.Println()
	bt.Root.PostPrintWithStack1()
	fmt.Println()
	fmt.Println("深度优先遍历: ")
	bt.Root.BreadthFirstPrint()
	fmt.Println()
	fmt.Println("层次遍历: ")
	bt.Root.LevelPrint()
	fmt.Println()
	fmt.Println("bt的高度: ", bt.Root.Hight(), ";bt的大小: ", bt.Root.Size())
}
