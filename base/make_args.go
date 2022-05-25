package base

import (
	"math/rand"
	"time"
)

func ArrGenerator(maxLen, minLen, maxV, minV int) []int {
	rand.Seed(time.Now().UnixNano())
	l := rand.Intn(maxLen-minLen+1) + minLen
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = rand.Intn(maxV+minV-1) + minV
	}
	return arr
}

func RandIntGenerator(n int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(n)
}
