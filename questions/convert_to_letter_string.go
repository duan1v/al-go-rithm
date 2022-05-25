package questions

// @brief: 将A-Z对应到1-26;一串纯数字组成的字符串s转换成A-Z的字符有多少种方法?如:123=>ABC|ABW|KC
func convertToLetterString(s string) (n int) {
	return convertToLetterStringProcess(0, s)
}

func convertToLetterStringProcess(cur int, s string) (n int) {
	if cur == len(s) {
		return 1
	}
	// 这是静态类型,是字符,需要加 ' ;
	if s[cur] == '0' {
		return
	}
	n = convertToLetterStringProcess(cur+1, s)
	if cur < len(s)-1 && (s[cur]-'0')*10+s[cur+1]-'0' < 27 {
		n += convertToLetterStringProcess(cur+2, s)
	}
	return
}

func TestConvertToLetterString() (r int) {
	// r = convertToLetterString("2132082")
	r = convertToLetterString("11111")
	return
}

// 动态规划
func dpConvertToLetterString(s string) (n int) {
	l := len(s)
	dp := make([]int, l+1)
	dp[l] = 1
	for i := l - 1; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}
		dp[i] = dp[i+1]
		if i < l-1 && (s[i]-'0')*10+s[i+1]-'0' < 27 {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}

func TestConvertToLetterString1() (r int) {
	// r = dpConvertToLetterString("2132082")
	r = dpConvertToLetterString("11111")
	return
}
