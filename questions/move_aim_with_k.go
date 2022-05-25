package questions

// @brief: 一共给出1~N,N个位置,机器人由start位置移动K次到aim位置的总共方法
// 一.尝试递归
func moveAimWithK(n int, start int, k int, aim int) int {
	if !validateNSKA(n, start, k, aim) {
		return -1
	}
	return moveAimWithKProcess(start, k, n, aim)
}

// @field cur 当前位置
// @field rest 剩余步数
func moveAimWithKProcess(cur int, rest int, n int, aim int) (w int) {
	if cur == aim {
		w = 1
	}
	if rest == 0 {
		return
	}

	// 如果当前位置是第一个,那么只能前进至2,走完之后还剩rest-1步
	if cur == 1 {
		w = moveAimWithKProcess(cur+1, rest-1, n, aim)
	} else if cur == n {
		// 当前在最后一个位置
		w = moveAimWithKProcess(cur-1, rest-1, n, aim)
	} else {
		// 在中间位置
		w = moveAimWithKProcess(cur+1, rest-1, n, aim) + moveAimWithKProcess(cur-1, rest-1, n, aim)
	}
	return
}

func validateNSKA(n int, start int, k int, aim int) bool {
	return n > 1 && start <= n && aim <= n && start > 0 && aim > 0 && k > 0
}

func TestMoveAimWithK() (r int) {
	return moveAimWithK(25, 2, 20, 10)
}

// ========================================================================================

// @optimize: 用缓存表记录,在cur位置,剩余rest步,到aim位置的总方法数:M
// 依据:
// 1)递归中的动态变量是cur与rest
// 2)可以用不同路线移动(K-rest)步由start到cur位置,此时就可以重复利用上面的M,省去重新递归的时间
// 即子方案重复才有优化的空间,动态规划的必要
// @keys: 从上向下的动态规划|记忆搜索法
func moveAimWithK1(n, start, k, aim int) int {
	if !validateNSKA(n, start, k, aim) {
		return -1
	}
	// cur总共有1~N种选择;rest总共有0~K种选择;初始化为-1,表示未被缓存
	auxiliary := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		x := make([]int, k+1)
		for j := 0; j <= k; j++ {
			x[j] = -1
		}
		auxiliary[i] = x
	}
	return moveAimWithKProcess1(start, k, n, aim, auxiliary)
}

// @field cur 当前位置
// @field rest 剩余步数
func moveAimWithKProcess1(cur int, rest int, n int, aim int, auxiliary [][]int) (w int) {
	if auxiliary[cur][rest] != -1 {
		return auxiliary[cur][rest]
	}

	if cur == aim {
		w = 1
	}
	if rest == 0 {
		return
	}

	// 如果当前位置是第一个,那么只能前进至2,走完之后还剩rest-1步
	if cur == 1 {
		w = moveAimWithKProcess(cur+1, rest-1, n, aim)
	} else if cur == n {
		// 当前在最后一个位置
		w = moveAimWithKProcess(cur-1, rest-1, n, aim)
	} else {
		// 在中间位置
		w = moveAimWithKProcess(cur+1, rest-1, n, aim) + moveAimWithKProcess(cur-1, rest-1, n, aim)
	}
	auxiliary[cur][rest] = w
	return
}

func TestMoveAimWithK1() (r int) {
	return moveAimWithK1(25, 2, 20, 10)
}

// =========================================================================================

// 对上面的缓存表进行分析,发现规律,全部填好,直接读取
// @tip:
// 1)显然第0列,只有aim行为1,其他均为0;
// 2)由上面的递归,可以发现,本次f(cur,rest)是被下次决定的,分析可知,
//   1>第一行的y列,由第二行的y-1列决定
//   2>第N行的y列,由第N-1行的y-1列决定
//   3>第[2,N-1]中的x行的y列,由x+1行的y-1列与x-1行的y-1列共同决定
func moveAimWithK2(n, start, k, aim int) int {
	if !validateNSKA(n, start, k, aim) {
		return -1
	}
	// cur总共有1~N种选择;rest总共有0~K种选择
	auxiliary := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		x := make([]int, k+1)
		auxiliary[i] = x
	}
	// 初始化第0列
	auxiliary[aim][0] = 1
	for j := 1; j <= k; j++ {
		auxiliary[1][j] = auxiliary[2][j-1]
		for i := 2; i < n; i++ {
			auxiliary[i][j] = auxiliary[i+1][j-1] + auxiliary[i-1][j-1]
		}
		auxiliary[n][j] = auxiliary[n-1][j-1]
	}
	return auxiliary[start][k]
}

func TestMoveAimWithK2() (r int) {
	return moveAimWithK2(25, 2, 20, 10)
}
