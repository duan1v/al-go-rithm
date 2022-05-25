package base

// arr是有序整型数组
func BinarySearch(arr []int, dest int, highIndex int, lowIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	midIndex := (highIndex + lowIndex) / 2
	midValue := arr[midIndex]
	if midValue == dest {
		return midIndex
	}
	if midValue > dest {
		return BinarySearch(arr, dest, midIndex-1, lowIndex)
	}
	if midValue < dest {
		return BinarySearch(arr, dest, highIndex, midIndex+1)
	}
	return -1
}

// 有序数组符合>=n条件时最左的位置
func LeftMostIndexSmallest(arr []int, n int) (t int) {
	highIndex := len(arr) - 1
	lowIndex := 0
	t = -1
	for lowIndex <= highIndex {
		midIndex := (highIndex + lowIndex) / 2
		midValue := arr[midIndex]
		if lowIndex == highIndex {
			if midValue >= n {
				t = midIndex
			}
			return t
		}
		if midValue >= n {
			t = midIndex
			highIndex = midIndex - 1
		}
		if midValue < n {
			lowIndex = midIndex + 1
		}
	}
	return t
}

// 有序数组符合<=n条件时最右的位置
func RightMostIndexBigest(arr []int, n int) (t int) {
	highIndex := len(arr) - 1
	lowIndex := 0
	t = -1
	for lowIndex <= highIndex {
		midIndex := (highIndex + lowIndex) / 2
		midValue := arr[midIndex]
		if lowIndex == highIndex {
			if midValue <= n {
				t = midIndex
			}
			return t
		}
		if midValue <= n {
			t = midIndex
			lowIndex = midIndex + 1
		}
		if midValue > n {
			highIndex = midIndex - 1
		}
	}
	return t
}

// 数组arr,无序且相邻不等;局部最小:比前后两个数都小,首不比前,尾不比后;找出一个局部最小位置
func PartSmallest(arr []int) int {
	l := len(arr)
	if l < 2 {
		return 0
	}
	if arr[0] < arr[1] {
		return 0
	}
	if arr[l-1] < arr[l-2] {
		return l - 1
	}

	// 1.上面两个判断规定了数组大概是这样的:\.../;所以二分之后呈\./.\或者\.\./,应取(\./)这部分
	// 2.数组长度为2时,上面的判断也能解决
	highIndex := l - 1
	lowIndex := 0
	// 3.处理长度不小于3的情况
	for lowIndex < highIndex-1 {
		midIndex := (lowIndex + highIndex) / 2
		midValue := arr[midIndex]
		// 4.下面这个判断可以但没必要
		// if !(midIndex > 0 && midIndex < l-1) {
		// 	return -1
		// }
		prevValue := arr[midIndex-1]
		nextValue := arr[midIndex+1]
		if midValue < prevValue && midValue < nextValue {
			return midIndex
		}
		if midValue < nextValue {
			highIndex = midIndex
		} else {
			lowIndex = midIndex
		}
	}
	// 如果lowIndex与highIndex相邻了,还未返回
	if arr[lowIndex] > arr[highIndex] {
		return highIndex
	}
	return lowIndex
}
