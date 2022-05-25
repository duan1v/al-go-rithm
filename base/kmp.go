package base

import "fmt"

// @brief: KMP算法:查找模式串在主串所在位置的算法
// @description: 与字符串暴力比对的差别是,跳过主串当前位置之前的字符的挨个匹配;即跳过不必要的匹配.
// 那么哪些是没必要的匹配?
// 如:主串S,模式串P;当S[i]与P[0]对齐时,在P[4]处失配,即S[i+4]!=P[4];
// 1)如果说S[i+1]是有必要与P[0]比较的话,那么就是说,至少,没有与已知条件相违背;
// 当前已知条件是:S[i:i+3]==P[0:3];
// 在P向前移动一格后,S[i+1:i+3]这三个数与P[0:2]对应相等才是有价值的;
// 又因S[i+1:i+3]==P[1:3];所以必须是P[0:3]都是相等的情况下,这个比较才是必要的;
// 2)如果说S[i+2]是有必要与P[0]比较的话,同上可得,P0==P2 && P1==P3才是必要的;
// 3)如果说S[i+3]是有必要与P[0]比较的话,同上可得,P0==P3才是必要的;

// 那么模式串应该向前移动几格?next数组就是干这事的.
// next[x]=k就是根据当前比较是否有必要,反向推出在P[x+1]处失配时,
// S[i+((x+1)-k)]与P[0]进行比较才是有必要的,也即S[i+(x+1)]与P[k]进行比较

// 我们可以正向地去看next,是很容易理解的.
// 在模式串中,用前缀替换后缀,直接用失配位置与前缀后的那个字符比对;
// 一次性前进了(模式串中的失配位置-前缀个数);而非一格一格前进;
// next[x]=k中,k不仅是P[0:x]中,前缀的长度;也是前缀替换后缀之后,与主串直接比对的位置
type KmpStr struct {
	Content string
	Next    []int
}

func (ks *KmpStr) MakeNext() {
	c := ks.Content
	l := len(c)
	next := make([]int, l)
	current := 1
	now := 0
	for current < l {
		if c[current] == c[now] {
			now++
			next[current] = now
			current++
		} else if now != 0 {
			now = next[now]
		} else {
			next[current] = 0
			current++
		}
	}
	ks.Next = next
}

func KmpSearch(source string, match string) []int {
	ks := KmpStr{Content: match}
	ks.MakeNext()

	sl := len(source)
	scur, mcur, poses := 0, 0, make([]int, 0)
	for scur < sl {
		if match[mcur] == source[scur] {
			scur++
			mcur++
		} else if mcur != 0 {
			mcur = ks.Next[mcur-1]
		} else {
			scur++
		}
		if mcur == len(match) {
			poses = append(poses, scur-mcur)
			mcur = ks.Next[mcur-1]
		}
	}
	return poses
}

func TestKmp() {
	ks := KmpStr{Content: "hahahs"}
	ks.MakeNext()
	fmt.Println(ks.Next)
	source := "ababaabaabacdasabaabacs"
	match := "abaabac"
	pos := KmpSearch(source, match)
	fmt.Println(pos)
}
