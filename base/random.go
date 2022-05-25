package base

import (
	cr "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

func IsReturnMinus(ws *sync.WaitGroup, c int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Seed(time.Now().Unix())
	ws.Add(1)
	go func() {
		for i := 0; i < c; i++ {
			x := r.Int()
			// fmt.Println("Int :", x)
			if x < 0 {
				fmt.Println("Int 可以随机到负数.")
				break
			}
		}
		ws.Done()
	}()
	ws.Add(1)
	go func() {
		for i := 0; i < c; i++ {
			x := r.Int63()
			// fmt.Println("Int63 :", x)
			if x < 0 {
				fmt.Println("Int63 可以随机到负数.")
				break
			}
		}
		ws.Done()
	}()
}

func IsReturnM1(ws *sync.WaitGroup, c int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Seed(time.Now().Unix())
	// ws.Add(1)
	// go func() {
	// 	// 1.此函数并不安全，当c=2000时，发生index out of range [-1]
	// 	for i := 0; i < c; i++ {
	// 		x := r.Float32() //[0.0,1.0)
	// 		// fmt.Println("Float32 :", x)
	// 		if x >= 1 {
	// 			fmt.Println("Float32 可以随机到不小于1的数.")
	// 			break
	// 		}
	// 	}
	// 	ws.Done()
	// }()
	ws.Add(1)
	go func() {
		rand.Seed(time.Now().Unix())
		// 2.而这里的就不会
		for i := 0; i < c; i++ {
			x := rand.Float32() //[0.0,1.0)
			if x >= 1 {
				fmt.Println("Float32 可以随机到不小于1的数.")
				break
			}
		}
		ws.Done()
	}()
	ws.Add(1)
	go func() {
		for i := 0; i < c; i++ {
			x := r.Float64() //[0.0,1.0),与java的math.random()类似
			// fmt.Println("Float64 :", x)
			if x >= 1 {
				fmt.Println("Float64 可以随机到不小于1的数.")
				break
			}
		}
		ws.Done()
	}()
}

func TestRandom() {
	// 全局函数
	// 设置随机种子
	rand.Seed(time.Now().Unix())
	fmt.Println(
		rand.Int(),
		rand.Intn(100),
		rand.Int31(),
		rand.Int31n(150)-50, // 可以获取[-50,149]之间的整数
		rand.Int63(),
		rand.Int63n(100),
		rand.Float32(),
		rand.Float64()*7, // 可以强制转化为整型,获取[0,6]之间的整数
	)
}

func TestRandomWithRandStruct() {
	// 设置随机种子 有点像MD5的salt
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(
		r.Int(),
		r.Intn(100),
		r.Int31(),
		r.Int31n(150)-50,
		r.Int63(),
		r.Int63n(100),
		r.Float32(),
		r.Float64(),
	)
}

func TestCryptoRand() {
	dice, err := cr.Int(cr.Reader, big.NewInt(1000)) // [0,max)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dice.Int64())
}

func TestRandFairness() {
	t := 10000
	s := 3000
	c := 0
	for i := 0; i < t; i++ {
		dice, _ := cr.Int(cr.Reader, big.NewInt(int64(t)))
		if dice.Int64() < int64(s) {
			c++
		}
	}

	rand.Seed(time.Now().UnixNano())

	c1 := 0
	for i := 0; i < t; i++ {
		dice := rand.Intn(t)
		if dice < s {
			c1++
		}
	}

	c2 := 0
	for i := 0; i < t; i++ {
		dice := rand.Float64()
		if dice*float64(t) < float64(s) {
			c2++
		}
	}

	fmt.Printf("%.6f\n%.6f\n%.6f\n", float64(c)/float64(t), float64(c1)/float64(t), float64(c2)/float64(t))
}

// 向上/向下及四舍五入取整
func TestApproximation() {
	x, y := 6.9, 5.2
	fmt.Println(int(x), math.Ceil(y), math.Floor(x), math.Floor(x+0.5), math.Floor(y+0.5))
}

// 题目一、已知存在获得[a,b]之间的随机整数的函数f(),如何通过f()获取[c,d]之间随机整数的函数g()
// 比如a=1,b=5;c=2,d=7
// 先写个f()
func f() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(5) + 1
}

func testFairnessBase(f func() int, n string) {
	m := make(map[int]int)
	for i := 0; i < 10000; i++ {
		dice := f()
		m[dice]++
	}
	fmt.Printf("%s %#v\n", n, m)
}

func TestFairness() {
	testFairnessBase(f, "函数f()的结果:")
	testFairnessBase(g, "函数g()的结果:")
	testFairnessBase(l, "函数l()的结果:")
	testFairnessBase(h1, "函数h1()的结果:")
}

// 转化成可以公平获取0,1的函数
func f1() int {
	for {
		x := f()
		if x < 3 {
			return 0
		}
		if x > 3 {
			return 1
		}
	}
}

// 利用左移获取[c,d]的大致取数范围
// f2()可以公平的获取[0,7]
func f2() int {
	return f1()<<2 + f1()<<1 + f1()
}

// 将f2()中不属于[c,d]的数字排除
func g() int {
	x := f2()
	if x < 2 {
		return g()
	}
	return x
}

// 题目二、已知h()返回0的概率为p,返回1的概率为1-p;试通过h()获取[c,d]之间随机整数的函数l()
// 先写个h()
func h() int {
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() > 0.86 {
		return 1
	}
	return 0
}

// 关键还是构造一个等几率返回0,1的函数h1()
func h1() int {
	x := h()
	y := h()
	if x^y == 0 {
		return h1()
	}
	return x
}

func h2() int {
	return h1()<<2 + h1()<<1 + h1()
}

func l() int {
	x := h2()
	if x < 2 {
		return l()
	}
	return x
}

func TestRandomCollector(ws *sync.WaitGroup) {
	TestRandom()
	TestRandomWithRandStruct()
	IsReturnMinus(ws, 10000)
	IsReturnM1(ws, 10000)
	TestCryptoRand()
	TestRandFairness()
	TestApproximation()
	TestFairness()
}
