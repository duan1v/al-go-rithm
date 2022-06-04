package questions

import (
	"fmt"

	"duan1v/al-go-rithm/base"
)

// @url: https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/
// @brief: 获取链表倒数第K个节点

// @tips: 快慢指针
func kthTOLast(head *base.NodeSL, k int) interface{} {
	q := head
	s := head

	for k > 0 {
		q = q.NextNode
		k--
	}

	// 现在q的长度是原链表长度l-k;当q遍历完全后,也就是完成了l-k次遍历;
	// 因为s与q同步遍历,所以s也遍历了1-k次,所以s的尾巴节点还有l-(l-k)个,即k个,也就是倒数第K个节点
	for q != nil {
		q = q.NextNode
		s = s.NextNode
	}

	return s.Data
}

func TestKthToLast() {
	ns7 := base.NodeSL{Data: 9}
	ns6 := base.NodeSL{Data: 1, NextNode: &ns7}
	ns5 := base.NodeSL{Data: 8, NextNode: &ns6}
	ns4 := base.NodeSL{Data: 2, NextNode: &ns5}
	ns3 := base.NodeSL{Data: 5, NextNode: &ns4}
	ns2 := base.NodeSL{Data: 3, NextNode: &ns3}
	ns1 := base.NodeSL{Data: 4, NextNode: &ns2}

	x := kthTOLast(&ns1, 5)
	fmt.Println(x)
}
