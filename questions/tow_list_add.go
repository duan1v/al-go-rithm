package questions

import "duan1v/al-go-rithm/base"

// @url: https://leetcode.cn/problems/add-two-numbers/
// @brief: 已知两个单链表的节点,输出规则是:2->3->4 + 2->5->6->9 = 4->8->0->0->1

func twoListAdd(first *base.NodeSL, second *base.NodeSL) *base.NodeSL {
	if first == nil {
		return second
	}
	if second == nil {
		return first
	}

	guardNode := &base.NodeSL{}
	rn := guardNode
	carry := 0

	for first != nil || second != nil {
		ld, sd := 0, 0
		if second != nil {
			sd, _ = (second.Data).(int)
			second = second.NextNode
		}
		if first != nil {
			ld, _ = (first.Data).(int)
			first = first.NextNode
		}
		rd := ld + sd + carry
		carry = 0
		if rd > 9 {
			rd = rd - 10
			carry = 1
		}
		rn.NextNode = &base.NodeSL{Data: rd}
		rn = rn.NextNode
		if second == nil && first == nil && carry == 1 {
			rn.NextNode = &base.NodeSL{Data: 1}
		}
	}
	return guardNode.NextNode
}

func TestTwoListAdd() {
	ns3 := base.NodeSL{Data: 3}
	ns2 := base.NodeSL{Data: 2, NextNode: &ns3}
	ns1 := base.NodeSL{Data: 1, NextNode: &ns2}

	ns41 := base.NodeSL{Data: 6}
	ns31 := base.NodeSL{Data: 9, NextNode: &ns41}
	ns21 := base.NodeSL{Data: 9, NextNode: &ns31}
	ns11 := base.NodeSL{Data: 4, NextNode: &ns21}

	s := twoListAdd(&ns1, &ns11)
	s.PrintWithHead()
}
