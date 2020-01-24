// もっとスマートな解答あり
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func pow(a, n int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret *= a
		}
		a *= a
		n >>= 1
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := nextInts(n)

	ans := 0
	lf := lcmFactorFast(a)
	for i := range a {
		f := 1
		for x, y := range lf {
			f = (f * pow(x, y)) % a[i]
		}
		f--
		if f < 0 {
			f = a[i] - 1
		}
		ans += f
	}
	puts(ans)
	wt.Flush()
}

var (
	sc  = bufio.NewScanner(os.Stdin)
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	wt  = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func nextFloat64s(n int) []float64 {
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		slice[i] = nextFloat64()
	}
	return slice
}

func nextMega() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
