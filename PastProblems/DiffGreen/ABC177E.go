// 解説AC
// エラトステネスの篩ver.が遅いと勘違い

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

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

func max(nums ...int) int {
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
	mf := minFactors(max(a...))
	for i := range a {
		for j := a[i]; j > 1; j /= mf[j] {
			m[i][mf[j]]++
		}
	}
	return m
}

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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := nextInts(n)

	pairwise := true
	pfs := primeFactors(a)
	fac := map[int]int{}
	for i := 0; i < n; i++ {
		for k := range pfs[i] {
			pairwise = pairwise && fac[k] == 0
			fac[k]++
		}
	}
	if pairwise {
		puts("pairwise coprime")
		return
	}

	G := a[0]
	for i := 1; i < n; i++ {
		G = gcd(G, a[i])
	}
	if G == 1 {
		puts("setwise coprime")
		return
	}

	puts("not coprime")
}
