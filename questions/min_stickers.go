package questions

import (
	"duan1v/al-go-rithm/base"
	"fmt"
	"sort"
	"strings"
)

// @url: https://leetcode.cn/problems/stickers-to-spell-word/
// @brief: stickers = ["with","example","science"], target = "thehat";输出3;每张纸片无数张;求最小纸片数
func minStickers(stickers []string, target string) int {
	// 可以排除明显-1的case,减少无效分支树
	// 比如含有目标串中的各字母的数量均比其他纸条少,且其他纸条包含全部目标串字母
	// 再比如把必要的字符串放在最前面
	// 由于target的长度决定了纸张的最多张数,可以设置一个达不到的张数
	tl := len(target)
	ans := minStickersProcess(stickers, tl+1, target)
	if ans < tl+1 {
		return ans
	}
	return -1
}

func minStickersProcess(stickers []string, invalidCount int, target string) (n int) {
	// 这条方案可解决
	if len(target) == 0 {
		return 0
	}
	// 每次选择下张纸条都可以是len(stickers)中的一个
	// 根据第一次选择的纸条编号获取划分len(stickers)中方式
	n = invalidCount
	for _, sticker := range stickers {
		res := minStickersMinus(sticker, target)
		// 这张纸条有可能帮忙收集target上的字母吗?
		if len(res) != len(target) {
			// 下面是n与递归结果比较,而不是invalidCount,手误一处浪费不少时间,这种问题怎么排查呢?...
			// n = base.Min(invalidCount, minStickersProcess(stickers, invalidCount, res))
			n = base.Min(n, minStickersProcess(stickers, invalidCount, res))
		}
	}
	// 这条方案不可解决
	// 跟之前不一样的是多了一个后置判断
	// 这是追溯判断有无解决的依据
	if n < invalidCount {
		// 有效帮助完成后算纸
		n++
	}
	return
}

func minStickersMinus(sticker string, nTarget string) (res string) {
	scount := [26]int{}
	for _, v := range sticker {
		scount[v-'a']--
	}
	for _, v := range nTarget {
		scount[v-'a']++
	}
	for k, y := range scount {
		if y > 0 {
			x := make([]byte, y)
			for j := 0; j < y; j++ {
				x[j] = byte(k) + 'a'
			}
			res = fmt.Sprintf("%s%s", res, string(x))
		}
	}
	return
}

func TestMinStickers() {
	params := map[string][]string{
		// "basicbasic": {"notice", "possible"},
		"thehat": {"with", "example", "science"},
		// "atomher":       {"these", "guess", "about", "garden", "him"},
		// "separatewhich": {"a", "enemy", "material", "whose", "twenty", "describe", "magnet", "put", "hundred", "discuss"},
	}
	for k, v := range params {
		// r := minStickers(v, k)
		// r := dfMinSticker(v, k)
		// r := bfMinStickers(v, k)
		r := dfMinStickers1(v, k)
		fmt.Println(r)
	}
}

// 缓存表
// 深度搜索
func dfMinSticker(stickers []string, target string) int {
	scount := make([][26]int, len(stickers))
	for i, sticker := range stickers {
		for _, v := range sticker {
			scount[i][v-'a']++
		}
	}
	memo := make(map[string]int)
	tl := len(target)
	ans := dpMinStickerProcess(memo, scount, tl+1, target)
	if ans < tl+1 {
		return ans
	}
	return -1
}

func dpMinStickerProcess(memo map[string]int, scount [][26]int, invalidCount int, target string) (n int) {
	if len(target) == 0 {
		return 0
	}

	split := strings.Split(target, "")
	sort.Strings(split)
	st := strings.Join(split, "")

	m, ok := memo[st]
	if ok {
		return m
	}

	n = invalidCount

	tcount := [26]int{}
	for _, v := range target {
		tcount[v-'a']++
	}
	for _, sticker := range scount {
		if sticker[target[0]-'a'] > 0 {
			res := ""
			for i, t := range tcount {
				if t > 0 {
					y := t - sticker[i]
					if y > 0 {
						x := make([]byte, y)
						for j := 0; j < y; j++ {
							x[j] = byte(i) + 'a'
						}
						res = fmt.Sprintf("%s%s", res, string(x))
					}
				}
			}
			n = base.Min(n, dpMinStickerProcess(memo, scount, invalidCount, res))
		}
	}
	if n < invalidCount {
		// 有效帮助完成后算纸
		n++
	}
	// 更新备忘录并返回
	memo[st] = n
	return
}

// 执行用时：4 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.6 MB, 在所有 Go 提交中击败了99.56%的用户
func dfMinStickers1(stickers []string, target string) int {
	// 每张贴纸上可贴字母的个数,不可篡改
	count := make([][26]byte, len(stickers))
	// 目标串各字母在哪些贴纸上有,不可篡改
	has := [26][]byte{}
	for k, sticker := range stickers {
		for _, v := range sticker {
			if strings.ContainsRune(target, v) {
				count[k][v-'a']++
				has[v-'a'] = append(has[v-'a'], byte(k))
			}
		}
	}

	tl := len(target)
	// 完成目标串不可能超过目标串字母总数
	invalidCount := tl + 1
	// 表示1<<invalidCount种目标串被贴上纸条的状态
	// 根据题目要求设成byte,省点空间
	dp := make([]byte, 1<<tl)
	for i := range dp {
		dp[i] = byte(invalidCount)
	}
	// 表示目标串为空串,则需要0个纸条
	dp[0] = 0

	// 一步一步贴上纸条
	for i := 0; i < (1 << tl); i++ {
		// 表示这个状态尚未未完成,无法基于这个状态完成目标串
		if dp[i] == byte(invalidCount) {
			continue
		}

		cur := 0
		// 在当前状态的基础上,挑一个未被填上的字母
		for j := range target {
			if i&(1<<j) == 0 {
				cur = j
				break
			}
		}

		// 从有cur位置上的这个字母的纸条上找
		for _, k := range has[target[cur]-'a'] {
			next := i
			// 只需要取出当前贴纸即可
			sc := count[k]
			// 不能浪费这张贴纸上的其他字母,继续更新当前状态
			for l, t := range target {
				if next&(1<<l) != 0 {
					// 表明这个状态下,l所在位置已经被完成
					continue
				}
				if sc[t-'a'] > 0 {
					next |= (1 << l)
					sc[t-'a']--
				}
			}
			dp[next] = base.Min(dp[next], dp[i]+1)
		}
	}
	if dp[1<<tl-1] == byte(invalidCount) {
		return -1
	}
	return int(dp[1<<tl-1])
}

// 广度搜索
func bfMinStickers(stickers []string, target string) int {
	memo := make(map[string]bool)
	res, count := target, 0
	queue := []map[string]int{{target: 0}}

	for len(queue) > 0 && count <= len(target) {
		cur := queue[0]
		for k, v := range cur {
			res, count = k, v
		}
		queue = queue[1:]
		for _, sticker := range stickers {
			scount := [26]int{}
			for _, s := range sticker {
				scount[s-'a']--
			}
			for _, s := range res {
				scount[s-'a']++
			}
			newRes := ""
			for i, s := range scount {
				if s > 0 {
					x := make([]byte, s)
					for j := 0; j < s; j++ {
						x[j] = byte(i) + 'a'
					}
					newRes = fmt.Sprintf("%s%s", newRes, string(x))
				}
			}
			_, ok := memo[newRes]
			if ok {
				continue
			}
			if newRes == "" {
				return count + 1
			}
			memo[newRes] = true
			if len(newRes) != len(res) {
				queue = append(queue, map[string]int{newRes: count + 1})
			}
		}
	}

	return -1
}
