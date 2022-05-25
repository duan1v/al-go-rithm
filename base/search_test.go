package base

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr       []int
		dest      int
		highIndex int
		lowIndex  int
	}
	type test struct {
		name string
		args args
		want int
	}
	// TODO: Add test cases.
	tests := []test{}

	for i := 0; i < 100; i++ {
		arr := ArrGenerator(9, 2, 949, -50)
		InsertSort(arr)
		l := len(arr)
		x := RandIntGenerator(l)
		tests = append(tests, test{fmt.Sprintf("test%d", i), args{arr, arr[x], l - 1, 0}, x})
		y := RandIntGenerator(10)
		if x > 0 && arr[x-1]+y < arr[x] {
			tests = append(tests, test{fmt.Sprintf("test%d-nil", i), args{arr, arr[x-1] + y, l - 1, 0}, -1})
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.arr, tt.args.dest, tt.args.highIndex, tt.args.lowIndex); got != tt.want {
				t.Errorf("BinarySearch() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
