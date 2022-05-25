package questions

import (
	"base"
	"fmt"
)

// @url: https://leetcode.cn/problems/stickers-to-spell-word/
// @brief: stickers = ["with","example","science"], target = "thehat";输出3;每张纸片无数张;求最小纸片数
func minStickers(stickers []string, target string) int {
	// 可以排除明显-1的case,减少无效分支树
	// 比如含有目标串中的各字母的数量均比其他纸条少,且其他纸条包含全部目标串字母
	// 再比如把必要的字符串放在最前面
	// 由于target的长度决定了纸张的最多张数,可以设置一个达不到的张数
	inValideCount := len(target) + 1
	ans := minStickersProcess(stickers, inValideCount, target)
	if ans < inValideCount {
		return ans
	}
	return -1
}

func minStickersProcess(stickers []string, inValideCount int, target string) (n int) {
	// 这条方案可解决
	if len(target) == 0 {
		return
	}
	// 每次选择下张纸条都可以是len(stickers)中的一个
	// 根据第一次选择的纸条编号获取划分len(stickers)中方式
	n = inValideCount
	for _, sticker := range stickers {
		res := minStickersMinus(sticker, target)
		// 这张纸条有可能帮忙收集target上的字母吗?
		if len(res) != len(target) {
			n = base.Min(inValideCount, minStickersProcess(stickers, inValideCount, res))
		}
	}
	// 这条方案不可解决
	// 更之前不一样的是多了一个后置判断
	// 这是追溯判断有无解决的依据
	if n < inValideCount {
		// 有效帮助完成后算纸
		n++
	}

	return
}

func minStickersMinus(sticker string, nTarget string) (res string) {
	sc := minStickersCountLetterWithStr(sticker)
	nc := minStickersCountLetterWithStr(nTarget)
	for k, v := range nc {
		if v > 0 {
			sv := sc[k]
			y := v - sv
			if y > 0 {
				x := make([]byte, y)
				for j := 0; j < y; j++ {
					x[j] = byte(k) + 'a'
				}
				res = fmt.Sprintf("%s%s", res, string(x))
			}
		}
	}
	return
}

func minStickersCountLetterWithStr(str string) (scount [26]int) {
	for _, v := range str {
		scount[v-'a']++
	}
	return
}

func TestinStickers() {
	params := map[string][]string{
		// "basicbasic": {"notice", "possible"},
		"thehat": {"with", "example", "science"},
		// "atomher": {"these", "guess", "about", "garden", "him"},
	}
	for k, v := range params {
		r := minStickers(v, k)
		fmt.Println(r)
	}
}
