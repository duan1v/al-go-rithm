package questions

import (
	"fmt"
)

// @brief: 有货质量q []int,对应价值v []int,背包容量为capacity,求可背运的最大价值(每件货只能取一次)
// @tip: 依次考虑当前货物是否放进背包,对下次选择的影响;暴力递归就是一个二叉树
func loadBag(q []int, v []int, cap int) (mv int) {
	if q == nil || len(q) != len(v) || cap < 1 {
		return
	}
	mv = loadBagProcess(q, v, 0, cap)
	return
}

// @field index 考虑要不要当前货物的位置
// @field res 背包剩余容量
func loadBagProcess(q []int, v []int, index int, res int) (mv int) {
	// 可能存在质量为0的货物,所以0继续
	if res < 0 {
		return
	}
	// 为了考虑len(q)-1的可能,需要放index=len(q)进来,但不继续往下走
	if index == len(q) {
		return
	}

	ca := loadBagProcess(q, v, index+1, res)
	// 当前货物质量超过剩余容量,这次就要不了了,只能返回ca
	if q[index] > res {
		return ca
	}

	mv = v[index] + loadBagProcess(q, v, index+1, res-q[index])

	if ca > mv {
		mv = ca
	}
	return
}

func TestLoadBag() (r int) {
	// r = loadBag([]int{3, 2, 4, 7, 3, 1, 7}, []int{5, 6, 3, 19, 12, 4, 2}, 15)
	// r = loadBag([]int{3, 2, 4, 7}, []int{5, 6, 3, 19}, 11)
	r = loadBag([]int{3, 2, 40, 7, 3, 1, 7}, []int{5, 6, 300, 19, 12, 4, 2}, 15)
	return
}

// ==================================================================================

// 动态规划
func dpLoadBag(q []int, v []int, cap int) (mv int) {
	n := len(q)
	if q == nil || n != len(v) || cap < 1 {
		return
	}

	anxiliary := make([][]int, n+1)
	// ps: go初始化多维数组好烦啊
	anxiliary[n] = make([]int, cap+1)

	for j := n - 1; j >= 0; j-- {
		anxiliary[j] = make([]int, cap+1)
		for k := 0; k <= cap; k++ {
			p1 := anxiliary[j+1][k]
			p2 := 0
			if q[j] <= k {
				p2 = v[j] + anxiliary[j+1][k-q[j]]
			}

			mv = p1
			if p1 < p2 {
				mv = p2
			}
			anxiliary[j][k] = mv
		}
	}

	return
}

func TestLoadBag1() (r int) {
	// r = dpLoadBag([]int{3, 2, 4, 7, 3, 1, 7}, []int{5, 6, 3, 19, 12, 4, 2}, 15)
	// r = dpLoadBag([]int{3, 2, 4, 7}, []int{5, 6, 3, 19}, 11)
	r = dpLoadBag([]int{3, 2, 40, 7, 3, 1, 7}, []int{5, 6, 300, 19, 12, 4, 2}, 15)
	return
}

// @brief: 是否可以将数组nums分割成两个子集,使得两个子集的元素和相等.
// @url: https://leetcode.cn/problems/partition-equal-subset-sum/
// @summary: 将整个过程视为,把nums的元素依次向已存在目标上添加;
// dp的索引就是可能达到的目标,0不需要任何数字都可以达到这个目标,所以dp[0]为ture;
// 这样的思路就是将dp的索引从可能变成可以的过程;
// 是最符合正向逻辑的,但遇到了几处思维漏洞,需要逐步尝试.
func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	if max > target {
		return false
	}

	dp := make([]bool, target+1)
	dp[0] = true

	// 1.为了不循环计算参数的各值,参数必须放在外层循环
	for _, num := range nums {
		if dp[target] {
			return true
		}
		// 2.因为dp[i+num]对i+num后面的dp中的值是有副作用的
		// 比如num=2,dp=[1,0,0,0,0...];
		// 0+2=2=>2+2=4=>2+4=6,而2已经是被制造出来的,也就是说4需要两次使用num,6要3次
		// 所以选择倒序
		for i := target - num; i >= 0; i-- {
			// 3.不过倒序会把之前的给覆盖,比如nums=[23, 13...]
			// 第一次循环:dp[23]=true;第二次循环dp[23]=dp[10+13]=dp[10]=false;这是不正确的
			// 所以说之前已经是可以达到的目标的话,接着为true,
			// 所以不可以写成下面的形式,我一开始想着要写成下面这种的
			// dp[i+num] = dp[i]
			// 所以取或
			dp[i+num] = dp[i] || dp[i+num]
		}
	}
	return dp[target]
}

func TestLoadBag2[T []int]() {
	params := []T{
		[]int{2, 2, 3, 5},
		[]int{3, 3, 3, 4, 5},
		[]int{1, 5, 5, 11},
		[]int{14, 9, 8, 4, 3, 2},
		[]int{23, 13, 11, 7, 6, 5, 5},
		[]int{1, 2, 3, 4, 5, 6, 7},
	}
	for _, v := range params {
		r := canPartition(v)
		fmt.Println(r)
	}
	// r := canPartition()
	// r := canPartition([]int{3, 3, 3, 4, 5})
	// r := canPartition([]int{1, 5, 5, 11})
	// r := canPartition([]int{14, 9, 8, 4, 3, 2})
	// r := canPartition([]int{23, 13, 11, 7, 6, 5, 5})
	// r := canPartition([]int{1, 2, 3, 4, 5, 6, 7})
}
