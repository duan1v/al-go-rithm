package questions

import (
	"duan1v/al-go-rithm/base"
)

// @url: https://leetcode.cn/problems/reverse-nodes-in-k-group/
// @brief: 已知链表头,要求k个一组,仅reverse组中链表,不够k的不动
func reverseKGroup(head *base.NodeSL, k int) *base.NodeSL {
	newHead := findKGroupEnd(head, k)
	if newHead == nil || k <= 1 {
		return head
	}

	oldStart := head
	start := newHead.NextNode
	reverseKGroupMeta(head, newHead, k)
	for start != nil {
		end := findKGroupEnd(start, k)
		if end == nil {
			break
		}
		reverseKGroupMeta(start, end, k)
		oldStart.NextNode = end
		oldStart = start
		start = start.NextNode
	}
	return newHead
}

// 使用哨兵节,确实可以提高代码的逻辑性和可读性,看起来更优雅。
func reverseKGroup1(head *base.NodeSL, k int) *base.NodeSL {
	firstKEnd := findKGroupEnd(head, k)
	if firstKEnd == nil || k <= 1 {
		return head
	}

	guardNode := &base.NodeSL{NextNode: head}
	prevStart := guardNode
	start := prevStart.NextNode
	for start != nil {
		end := findKGroupEnd(start, k)
		if end == nil {
			break
		}
		reverseKGroupMeta(start, end, k)
		prevStart.NextNode = end
		prevStart = start
		start = start.NextNode
	}
	return guardNode.NextNode
}

func findKGroupEnd(start *base.NodeSL, k int) *base.NodeSL {
	for ; start != nil && k > 1; k-- {
		start = start.NextNode
	}
	return start
}

func reverseKGroupMeta(start *base.NodeSL, end *base.NodeSL, k int) {
	pn := end.NextNode
	for ; k > 0; k-- {
		nn := start.NextNode
		start.NextNode = pn
		pn = start
		start = nn
	}
}

func TestReverseKGroup() {
	ns6 := base.NodeSL{Data: 6}
	ns5 := base.NodeSL{Data: 5, NextNode: &ns6}
	ns4 := base.NodeSL{Data: 4, NextNode: &ns5}
	ns3 := base.NodeSL{Data: 3, NextNode: &ns4}
	ns2 := base.NodeSL{Data: 2, NextNode: &ns3}
	ns1 := base.NodeSL{Data: 1, NextNode: &ns2}
	start := reverseKGroup(&ns1, 5)
	start.PrintWithHead()
	start = reverseKGroup1(start, 2)
	start.PrintWithHead()
}
