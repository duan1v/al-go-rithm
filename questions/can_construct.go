package questions

import "fmt"

// @url: https://leetcode.cn/problems/ransom-note/
// @brief: 字符串a是否可以由给定字符串b中的字母获得
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	scount := [26]int32{}

	for _, m := range magazine {
		scount[m-'a']++
	}

	for _, r := range ransomNote {
		scount[r-'a']--
		if scount[r-'a'] < 0 {
			return false
		}
	}
	return true
}

// 下面这种写法循环数字好像更快点
func canConstruct1(ransomNote string, magazine string) bool {
	lr, lm := len(ransomNote), len(magazine)
	if lr > lm {
		return false
	}
	scount := [26]int32{}
	for i := 0; i < lm; i++ {
		scount[magazine[i]-'a']++
	}
	for i := 0; i < lr; i++ {
		ri := ransomNote[i] - 'a'
		if scount[ransomNote[i]-'a'] <= 0 {
			return false
		}
		scount[ri]--
	}
	return true
}

func TestCanConstruct() {
	params := map[string]string{
		"aa": "aab",
	}
	for k, v := range params {
		r := canConstruct1(k, v)
		// r := canConstruct(k, v)
		fmt.Println(r)
	}

}
