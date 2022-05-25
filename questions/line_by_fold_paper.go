package questions

import (
	"base"
	"fmt"
	"strings"
)

// @brief: 将一张纸对折后打开,可得到一个凹痕,输出down,
// 再次对折打开,可顺序得到,凹凹凸,输出down down up;求n次对折的输出

// @tip: 实际折纸可发现,在折完n-1次后,在折第n次时,是将n-1次的折痕翻上去,
// 这样n-1折痕的前后就会出现一个凹折痕和一个凸折痕;所以可以看作是生成子节点,
// 那么此题就可以看成以第一个凹痕为根节点,高度为n的满二叉树的中序打印

func lineByFoldPaper(n int) {
	// 创建树
	x := &base.NodeBT{Data: "down"}
	createFullBinaryTree(x, n)
	x.InPrint()
}

// 根据高度生成满二叉树
func createFullBinaryTree(nb *base.NodeBT, h int) {
	if h > 1 {
		nb.LeftNode = &base.NodeBT{Data: "down"}
		nb.RightNode = &base.NodeBT{Data: "up"}
		createFullBinaryTree(nb.LeftNode, h-1)
		createFullBinaryTree(nb.RightNode, h-1)
	}
}

// 不生成树,只利用中序的思路
func lineByFoldPaperWithoutTree(n int) {
	lineByFoldPaperWithoutTreeProcess(1, false, n)
}

func lineByFoldPaperWithoutTreeProcess(cur int, isUp bool, n int) {
	if cur > n {
		return
	}
	x := "down"
	if isUp {
		x = "up"
	}
	lineByFoldPaperWithoutTreeProcess(cur+1, false, n)
	fmt.Printf("%s ", x)
	lineByFoldPaperWithoutTreeProcess(cur+1, true, n)
}

func TestLineByFoldPaper() {
	lineByFoldPaper(1)
	fmt.Println()
	lineByFoldPaper(2)
	fmt.Println()
	lineByFoldPaper(3)
	fmt.Println()
	lineByFoldPaper(4)
	fmt.Println()
	fmt.Println(strings.Repeat("=", 80))
	lineByFoldPaperWithoutTree(1)
	fmt.Println()
	lineByFoldPaperWithoutTree(2)
	fmt.Println()
	lineByFoldPaperWithoutTree(3)
	fmt.Println()
	lineByFoldPaperWithoutTree(4)
}
