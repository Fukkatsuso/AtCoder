package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(nums []int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

// 素因数分解
// 自然数nを素因数分解して
// mapで返す
func primeFactor(n int) map[int]int {
	m := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n] = 1
	}
	return m
}

// 最大公約数
// 素数の積で表して返す
func gcdFactor(a []int) map[int]int {
	m := primeFactor(a[0])
	for i := 1; i < len(a); i++ {
		pf := primeFactor(a[i])
		for x := range m {
			if pf[x] == 0 {
				delete(m, x)
			} else if m[x] > pf[x] {
				m[x] = pf[x]
			}
		}
	}
	return m
}

// 最小公倍数
// 素数の積で表して返す
func lcmFactor(a []int) map[int]int {
	m := map[int]int{}
	for _, n := range a {
		pf := primeFactor(n)
		for x, y := range pf {
			if m[x] < y {
				m[x] = y
			}
		}
	}
	return m
}

// nまでの自然数の最小の素因数
func minFactors(n int) []int {
	mf := make([]int, n+1)
	mf[0], mf[1] = -1, -1
	for i := 2; i <= n; i++ {
		for j := 1; i*j <= n; j++ {
			if mf[i*j] == 0 {
				mf[i*j] = i
			}
		}
	}
	return mf
}

// 高速素因数分解
func primeFactors(a []int) []map[int]int {
	// 各数の素因数分解結果を入れるmap配列
	m := make([]map[int]int, len(a))
	for i := range m {
		m[i] = map[int]int{}
	}
	// 最小素因数テーブル
	mf := minFactors(max(a))
	for i := range a {
		for j := a[i]; j > 1; j /= mf[j] {
			m[i][mf[j]]++
		}
	}
	return m
}

// 最大公約数の高速版
// 素因数分解した形で返す
func gcdFactorFast(a []int) map[int]int {
	pfs := primeFactors(a)
	m := pfs[0]
	for i := 1; i < len(a); i++ {
		for x := range m {
			if m[x] > pfs[i][x] {
				m[x] = pfs[i][x]
			}
			if m[x] == 0 {
				delete(m, x)
			}
		}
	}
	return m
}

// 最小公倍数
// 素因数分解した形で返す
func lcmFactorFast(a []int) map[int]int {
	pfs := primeFactors(a)
	m := map[int]int{}
	for x, y := range pfs[0] {
		m[x] = y
	}
	for i := 1; i < len(a); i++ {
		for x := range pfs[i] {
			if pfs[i][x] > m[x] {
				m[x] = pfs[i][x]
			}
		}
	}
	return m
}

func main() {
	// puts(primeFactor(60))
	puts(primeFactors([]int{12, 17, 20, 36, 60}))
	puts(gcdFactor([]int{12, 17, 20, 36, 60}))
	// puts(lcmFactor([]int{2, 3, 5, 4}))
	// puts(minFactors(20))
	puts(gcdFactorFast([]int{12, 20, 36, 60}))
	puts(lcmFactorFast([]int{12, 20, 36, 60}))
	wt.Flush()
}

var (
	wt = bufio.NewWriter(os.Stdout)
)

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
