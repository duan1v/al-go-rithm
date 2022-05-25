package questions

import "base"

// @url: https://leetcode.cn/problems/merge-two-sorted-lists/
// @brief: 输入：l1 = [1,2,4], l2 = [1,3,4] 输出：[1,1,2,3,4,4]

func mergeTwoLists(list1 *base.NodeSL, list2 *base.NodeSL) *base.NodeSL {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	guardNode := &base.NodeSL{}
	rn := guardNode
	cur1 := list1
	cur2 := list2

	for cur1 != nil && cur2 != nil {
		cd1, _ := cur1.Data.(int)
		cd2, _ := cur2.Data.(int)
		if cd1 >= cd2 {
			rn.NextNode = cur2
			cur2 = cur2.NextNode
		} else {
			rn.NextNode = cur1
			cur1 = cur1.NextNode
		}
		rn = rn.NextNode
		if cur1 == nil {
			rn.NextNode = cur2
		}
		if cur2 == nil {
			rn.NextNode = cur1
		}
	}
	return guardNode.NextNode
}

func TestMergeTwoList() {
	ns3 := base.NodeSL{Data: 9}
	ns2 := base.NodeSL{Data: 8, NextNode: &ns3}
	ns1 := base.NodeSL{Data: 4, NextNode: &ns2}

	ns41 := base.NodeSL{Data: 11}
	ns31 := base.NodeSL{Data: 7, NextNode: &ns41}
	ns21 := base.NodeSL{Data: 3, NextNode: &ns31}
	ns11 := base.NodeSL{Data: 1, NextNode: &ns21}

	s := mergeTwoLists(&ns1, &ns11)
	s.PrintWithHead()
}
