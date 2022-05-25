package questions

// @brief: 一组牌,每张牌都有一个分数,两人取牌,只能取剩余牌中第一张或最后一张;
// 取完时,分数大的获胜;两人都想获胜的心态,求最后获胜的分数

// 策略:使下一次拿的人获取更少的分
func cardObtainScoreMax(cards []int) (first int) {
	l := len(cards)
	if cards == nil || l == 0 {
		return 0
	}
	first = cardObtainScoreMaxFrontProcess(0, l-1, cards)
	second := cardObtainScoreMaxBehindProcess(0, l-1, cards)
	if second > first {
		first = second
	}
	return
}

// 先手函数
func cardObtainScoreMaxFrontProcess(l int, r int, cards []int) (s int) {
	if l == r {
		return cards[l]
	}
	sl := cards[l] + cardObtainScoreMaxBehindProcess(l+1, r, cards)
	s = cards[r] + cardObtainScoreMaxBehindProcess(l, r-1, cards)
	if sl > s {
		s = sl
	}
	return
}

// 后手函数
func cardObtainScoreMaxBehindProcess(l int, r int, cards []int) (s int) {
	if l == r {
		return 0
	}
	sl := cardObtainScoreMaxFrontProcess(l+1, r, cards)
	s = cardObtainScoreMaxFrontProcess(l, r-1, cards)
	if sl < s {
		s = sl
	}
	return
}

func TestCardObtainScoreMax() (r int) {
	return cardObtainScoreMax([]int{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7})
}

// =========================================================================

// 生成缓存表;依据递归的动态变量l,r;先手函数,后手函数各自生成缓存表

// 生成递归结束条件生成最基础数据,根据递归规则,及基础数据完成缓存表
func cardObtainScoreMax1(cards []int) (first int) {
	n := len(cards)
	af := make([][]int, n)
	ab := make([][]int, n)
	// 以当前第一个元素的位置l为纵坐标;最后一个元素的位置r为横坐标
	// 初始化两张表及基础数据
	for i := 0; i < n; i++ {
		y := make([]int, n)
		ab[i] = make([]int, n)
		y[i] = cards[i]
		af[i] = y
	}

	// 关于对角线:
	// 遍历一共有多少列对角线
	for j := 1; j < n; j++ {
		// 遍历当前这一列对角线中的每一格
		// 前后手缓存可以写在同一个循环中,因为各自都有第0列的对角线作为基础数据
		for k := j; k < n; k++ {
			fg := ab[k-j][k-1] + cards[k]
			if fg < ab[k-j+1][k]+cards[k-j] {
				fg = ab[k-j+1][k] + cards[k-j]
			}
			af[k-j][k] = fg

			bg := af[k-j][k-1]
			if bg > af[k-j+1][k] {
				bg = af[k-j+1][k]
			}
			ab[k-j][k] = bg
		}

		// 或者根据当前列对角线初始坐标(0,j)
		// k := 0
		// l := j
		// for l < n {
		// 	fg := ab[k][l-1] + cards[l]
		// 	if fg < ab[k+1][l]+cards[k] {
		// 		fg = ab[k+1][l] + cards[k]
		// 	}
		// 	af[k][l] = fg

		// 	bg := af[k][l-1]
		// 	if bg > af[k+1][l] {
		// 		bg = af[k+1][l]
		// 	}
		// 	ab[k][l] = bg
		// 	k++
		// 	l++
		// }
	}

	first = af[0][n-1]
	second := ab[0][n-1]
	if second > first {
		first = second
	}
	return
}

func TestCardObtainScoreMax1() (r int) {
	return cardObtainScoreMax1([]int{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7})
}
