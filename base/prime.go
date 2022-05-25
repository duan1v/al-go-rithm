package base

func RudeOfSumPrime(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime {
			count++
		}
	}
	return count
}

// 埃筛法查找质数
func Eratosthenes(n int) int {
	isNotPrime := make([]bool, n+1)
	count := 0
	for i := 2; i <= n; i++ {
		if !isNotPrime[i] {
			count++
			for j := i * i; j <= n; j += i {
				isNotPrime[j] = true
			}
		}
	}
	return count
}
