package questions

import "fmt"

// @brief: 给定数组arr及其中两个索引L,R(L<=R);求取arr在L~R位置的和(包含L,R).

// 制作辅助表,因为是直接返回值，所以适合大数量级的循环取和
func SumArrWithIndexesByMap(arr []int, l int, r int) int {
	m := make(map[string]int)
	for k, v := range arr {
		m[fmt.Sprintf("%d-%d", k, k)] = v
		for k1, v1 := range arr {
			if k1 <= k {
				continue
			}
			m[fmt.Sprintf("%d-%d", k, k1)] = m[fmt.Sprintf("%d-%d", k, k1-1)] + v1
		}
	}

	return m[fmt.Sprintf("%d-%d", l, r)]
}

// 制作辅助数组h,h在索引x位置的值为,arr在x位置至0位置的和;所以所求为h在r位置的值-(l-1)位置的和
func SumArrWithIndexesByArr(arr []int, l int, r int) int {
	h := make([]int, len(arr))
	for k, v := range arr {
		if k == 0 {
			h[k] = v
			continue
		}
		h[k] = h[k-1] + v
	}
	if l == 0 {
		return h[r]
	}
	return h[r] - h[l-1]
}
