package base

func validate(arr []int) {
	if len(arr) < 2 {
		panic("待排序数组长度必须大于1!")
	}
}

func SelectSort(arr []int) {
	validate(arr)
	al := len(arr)
	for i := 0; i < al-1; i++ {
		minIndex := i
		for j := i + 1; j < al; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func BubbleSort(arr []int) {
	validate(arr)
	al := len(arr)
	for i := al - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func InsertSort(arr []int) {
	validate(arr)
	al := len(arr)
	for i := 1; i < al; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}
