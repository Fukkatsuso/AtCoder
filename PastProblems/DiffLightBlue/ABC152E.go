package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	const mod = 1000000007
	n := nextInt()
	a := nextInts(n)
	// 各aの素因数分解
	pfs := primeFactors(a)
	// 最小公倍数の素因数分解
	lcmp := lcmPrimes(pfs)

	// 最小公倍数
	lcm := 1
	for x, y := range lcmp {
		for i := 0; i < y; i++ {
			lcm = (lcm * x) % mod
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		b := lcm * modPow(a[i], mod-2, mod)
		ans = (ans + b) % mod
	}
	puts(ans)
	wt.Flush()
}

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func max(nums []int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
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

func lcmPrimes(pfs []map[int]int) map[int]int {
	m := map[int]int{}
	for x, y := range pfs[0] {
		m[x] = y
	}
	for i := 1; i < len(pfs); i++ {
		for x := range pfs[i] {
			if pfs[i][x] > m[x] {
				m[x] = pfs[i][x]
			}
		}
	}
	return m
}

// (a^n) % mod
func modPow(a, n, mod int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret = (ret * a) % mod
		}
		a = (a * a) % mod
		n >>= 1
	}
	return ret
}
