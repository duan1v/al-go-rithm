package questions

import (
	"base"
	"fmt"
)

// @url: https://leetcode.cn/problems/longest-common-subsequence/
// @brief: 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)

	return longestCommonSubsequenceProcess(text1, text2, l1-1, l2-1)
}

// 从右到左读s1,s2;去拼接公共子序列
func longestCommonSubsequenceProcess(s1 string, s2 string, si int, sj int) (p int) {
	// 无字符可继续比较则返回0
	if si < 0 || sj < 0 {
		return
	}

	if s1[si] == s2[sj] {
		// 当前位置的字符相同时,判断下一个字符
		p = 1 + longestCommonSubsequenceProcess(s1, s2, si-1, sj-1)
	} else {
		// 否则有两种后续方式:s1[si] vs s2[sj-1];s2[sj] vs s1[si-1]
		// 取大
		p = longestCommonSubsequenceProcess(s1, s2, si, sj-1)
		p2 := longestCommonSubsequenceProcess(s1, s2, si-1, sj)
		if p < p2 {
			p = p2
		}
	}

	return
}

// 缓存表
func dpLongestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	dp := make([][]int, l1+1)
	dp[0] = make([]int, l2+1)
	for i := 1; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				// 当前位置的字符相同时,判断下一个字符
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				// 否则有两种后续方式:s1[si] vs s2[sj-1];s2[sj] vs s1[si-1]
				// 取大
				dp[i][j] = dp[i][j-1]
				p := dp[i-1][j]
				if dp[i][j] < p {
					dp[i][j] = p
				}
			}
		}
	}

	return dp[l1][l2]
}

// @tip: 用text2的所有字符与text1的每个字符比较
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了100.00%的用户
func longestCommonSubsequence2(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	// 根据题目限制,用数组节省空间
	dp := [1001]int{}
	dp1 := [1001]int{}
	// text1与text2一端对齐,这层循环每一次都是将text1当作text1[:i+2],去与text2比较
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			// 必须是在上一次完成了之前的目标串text1[:i]的基础上
			if text1[i-1] == text2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				// 这一步只是通知下dp[j],上一次循环text1[i-1]已经被text2[j-1]匹配了,子序列长度增加了
				dp[j] = base.Max(dp[j], dp[j-1])
			}
			// 如:text1:ziz;text2:xzx;那么text1[0]与text2[0]匹配到之后,
			// text1进入到下一轮循环之后,gp=[0,0,1,1];gp1=[0,0,1,1]
			// 由gp1[1]与gp1[2]比较可以看出,
			// text2匹配完成text1本轮之前的循环,text2需要派出1个字符,
			// gp1保留了text2完成text1的节点记录:text2的位置1;
			// text2开始循环匹配text1[1],假设匹配成功,
			// 在上面的判断(text1[i-1] == text2[j-1])中,是否需要加1?
			// 1)如果是在上一次匹配成功的位置之前,是需要加1的,这表示一个不同于上一次的子序列的开头;
			// 2)如果是在上一次匹配成功的位置之后,是需要加1的,这表示这个位置之前的位置,
			// text2已经派出足够的字母去匹配成功了;
			// 3)如果是在上一次匹配成功的位置上呢?如果加的话,在下一次的循环中就很明显了,
			// text2[1]会继续匹配text1[2],而且匹配成功,这就会造成重复匹配,
			// 所以显然,不能直接使用gp[j]++
			// 那么如果就是要用gp[j]++,可行吗?
			// 1)好像使判断中 && dp[j] > dp[j-1] 不就行了?这样不就找到转折点了吗?
			// 显然过于严格,第一个字母匹配都加不了
			// 2)又或者gp[j]=gp[j-1]+1不就行了吗?重设text1,text2:="xzx","ziz"
			// 这个会造成text2本轮循环直接重复
			// 3)那么直接1个flag,标识当前位置是否是上一次匹配成功的位置不就行了吗?
			// 答案是也不行!比如：text1,text2:="zizbhb","xzxaba";对于这种就需要两个flag,
			// 还得知道在哪里flag,可以说很难操作
			// 综上,最长子序列研究结束,爽!
			dp1[j] = dp[j-1]
		}
	}
	return dp[l2]
}

// 这是LeetCode上原来的内存最少范例
// 还是没有完全懂这个x的用法
func longestCommonSubsequence4(text1 string, text2 string) int {
	x := 0
	dp := make([]int, len(text2)+1)
	for i := 0; i <= len(text1); i++ {
		for j := 0; j <= len(text2); j++ {
			tmp := dp[j]
			if i == 0 || j == 0 {
				dp[j] = 0
			} else if text1[i-1] == text2[j-1] {
				dp[j] = x + 1
			} else {
				dp[j] = base.Max(dp[j], dp[j-1])
			}
			x = tmp
		}
	}
	return dp[len(text2)]
}

// 这是我根据上面那个方法改的
// 我只知道x是记录上一次对text2循环中text2完成text1中字段的状态
// 但还没完全想通x具体解决了哪些思维漏洞,相比于直接d[j]++
func longestCommonSubsequence3(text1 string, text2 string) int {
	x := 0
	l1, l2 := len(text1), len(text2)
	dp := [1001]int{}
	for i := 1; i <= l1; i++ {
		x = 0
		for j := 1; j <= l2; j++ {
			temp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = x + 1
			} else {
				dp[j] = base.Max(dp[j], dp[j-1])
			}
			x = temp
		}
	}
	return dp[l2]
}

func TestLongestCommonSubsequence() {
	params := map[string]string{
		// "abcde":      "ace",
		// "aa":         "aab",
		// "oxcpqrsvwf": "shmtulqrypy",
		// "zizz": "xzzx",
		// "a": "aa",
		// "vcnwrmxc":   "pmlstotylonkvmhqjyxmnq",
		// "bmvcnwrmxcfcxabkxcvgbozmpspsbenazglyxkpibgzq": "bmpmlstotylonkvmhqjyxmnqzctonqtobahcrcbibgzgx",
		// "mhziwb":   "mhzziwb",
		// "bsb": "b",
		// "mhunuzqrkzsnidwbun": "szulspmhwpazoxijwbq",
		// "yzyn":  "zxyzm",
		// "yy":    "zxyzm",
		// "abcba":  "abcbcb",
		// "abcxba": "abcxbcxb",
		// "bcb": "bcbcb",
		// "ziz": "xzx",
		// "xzx":    "ziz",
		"zizbhb": "xzxaba",
	}
	for k, v := range params {
		// r := longestCommonSubsequence(k, v)
		// r := dpLongestCommonSubsequence(k, v)
		r := longestCommonSubsequence2(k, v)
		// r := longestCommonSubsequence3(k, v)
		fmt.Println(r)
	}

}
