package main

import (
	"fmt"
	"questions"
	"sync"
	"time"
)

//@brief: 耗时统计函数
func timeCost() func(flag string) {
	start := time.Now()
	return func(flag string) {
		tc := time.Since(start)
		fmt.Printf("%stime cost = %v\n", flag, tc)
	}
}

func wrapTestCase(ws *sync.WaitGroup, flag string, count int, function func()) {
	ws.Add(1)
	go func() {
		defer timeCost()(flag)
		for i := 0; i < count; i++ {
			function()
		}
		ws.Done()
	}()
}

func main() {
	ws := sync.WaitGroup{}
	defer ws.Wait()

	// 测试随机数
	// base.TestRandomCollector(&ws)
	// base.TestSingleList()
	// base.TestDoubleList()
	// base.TestQueue()
	// base.TestStack()
	// base.TestDoubleQueue()
	// base.TestKmp()
	// base.TestBinaryTree()
	// base.TestSQueue()
	// base.TestSStack()

	// 可以同时打开
	// wrapTestCase(&ws, "RudeOfSumPrime ", 100, func() {
	// 	base.RudeOfSumPrime(1000)
	// })
	// wrapTestCase(&ws, "Eratosthenes ", 100, func() {
	// 	base.Eratosthenes(1000)
	// })
	// wrapTestCase(&ws, "Fibonacci ", 2, func() {
	// 	base.Fibonacci(30)
	// })
	// wrapTestCase(&ws, "SelectSort ", 200, func() {
	// 	arr := []int{321, 543, 65, 47, 7869, 780, 908, 312}
	// 	base.SelectSort(arr)
	// 	// fmt.Printf("arr:%#v\n", arr)
	// })
	// wrapTestCase(&ws, "BubbleSort ", 200, func() {
	// 	arr := []int{321, 543, 65, 47, 7869, 780, 908, 312}
	// 	base.BubbleSort(arr)
	// 	// fmt.Printf("arr:%#v\n", arr)
	// })
	// wrapTestCase(&ws, "InsertSort ", 200, func() {
	// 	arr := []int{321, 543, 65, 47, 7869, 780, 908, 312}
	// 	base.InsertSort(arr)
	// 	// fmt.Printf("arr:%#v\n", arr)
	// })
	// wrapTestCase(&ws, "SumArrWithIndexesByMap ", 1, func() {
	// 	arr := []int{321, 543, 65, 47, 7869, 780, 908, 312}
	// 	x := questions.SumArrWithIndexesByMap(arr, 1, 3)
	// 	fmt.Printf("x:%#v\n", x)
	// })
	// wrapTestCase(&ws, "SumArrWithIndexesByArr ", 1, func() {
	// 	arr := []int{321, 543, 65, 47, 7869, 780, 908, 312}
	// 	x := questions.SumArrWithIndexesByArr(arr, 1, 3)
	// 	fmt.Printf("x:%#v\n", x)
	// })

	// arr := []int{321, 54, 65, 47, 7869, 780, 908, 950}
	// base.InsertSort(arr)
	// // al := len(arr)
	// // highIndex := al - 1
	// // lowIndex := 0
	// // wrapTestCase(&ws, "BinarySearch ", 1, func() {
	// // 	x := base.BinarySearch(arr, 321, highIndex, lowIndex)
	// // 	fmt.Println(arr, x)
	// // })
	// wrapTestCase(&ws, "LeftMostIndexSmallest ", 1, func() {
	// 	x := base.LeftMostIndexSmallest(arr, 322)
	// 	fmt.Println(arr, x)
	// })
	// wrapTestCase(&ws, "RightMostIndexBigest ", 1, func() {
	// 	x := base.RightMostIndexBigest(arr, 323)
	// 	fmt.Println(arr, x)
	// })
	// arr1 := []int{3, 2, 3, 2, 1, 3}
	// wrapTestCase(&ws, "PartSmallest ", 1, func() {
	// 	x := base.PartSmallest(arr1)
	// 	fmt.Println(arr1, x)
	// })
	// questions.TestReverseKGroup()
	// questions.TestTwoListAdd()
	// questions.TestMergeTwoList()
	// questions.TestKthToLast()

	// wrapTestCase(&ws, "TestMoveAimWithK ", 100, func() {
	// 	questions.TestMoveAimWithK()
	// 	// fmt.Println("总共方法为: ", r)
	// })
	// wrapTestCase(&ws, "TestMoveAimWithK1 ", 100, func() {
	// 	// 怎么优化过的还没上面的快...
	// 	questions.TestMoveAimWithK1()
	// 	// fmt.Println("总共方法为: ", r)
	// })
	// wrapTestCase(&ws, "TestMoveAimWithK2 ", 100, func() {
	// 	questions.TestMoveAimWithK2()
	// 	// fmt.Println("总共方法为: ", r)
	// })
	// questions.TestLineByFoldPaper()
	// wrapTestCase(&ws, "TestCardObtainScoreMax ", 1, func() {
	// 	r := questions.TestCardObtainScoreMax()
	// 	fmt.Println("最大分数为: ", r)
	// })
	// wrapTestCase(&ws, "TestCardObtainScoreMax1 ", 1, func() {
	// 	r := questions.TestCardObtainScoreMax1()
	// 	fmt.Println("最大分数为: ", r)
	// })
	// wrapTestCase(&ws, "TestLoadBag ", 1000, func() {
	// 	questions.TestLoadBag()
	// 	// fmt.Println("最大价值为: ", r)
	// })
	// wrapTestCase(&ws, "TestLoadBag1 ", 1, func() {
	// 	questions.TestLoadBag1()
	// 	// fmt.Println("最大价值为: ", r)
	// })
	// questions.TestLoadBag2()
	// r := questions.TestConvertToLetterString()
	// fmt.Println(r)
	// r = questions.TestConvertToLetterString1()
	// fmt.Print(r)
	// questions.TestMinStickers()
	// questions.TestCanConstruct()
	questions.TestLongestCommonSubsequence()
}
