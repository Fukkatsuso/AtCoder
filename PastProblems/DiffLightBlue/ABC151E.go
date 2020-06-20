// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

const COMMAX int64 = 100001
const MOD int64 = 1000000007

var fac [COMMAX]int64
var finv [COMMAX]int64
var inv [COMMAX]int64

// テーブル作成
func COMinit() {
	fac[0] = 1
	fac[1] = 1
	finv[0] = 1
	finv[1] = 1
	inv[1] = 1
	for i := int64(2); i < COMMAX; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
}

// 二項係数nCk
func COM(n, k int64) int64 {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % MOD) % MOD
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	COMinit()

	n, k := int64(nextInt()), int64(nextInt())
	a := nextInts(int(n))
	sort.Sort(sort.IntSlice(a))

	ans := int64(0)
	for i := 0; i < int(n)-1; i++ {
		fac := COM(n, k)
		if fac -= COM(int64(i+1), k); fac < 0 {
			fac += MOD
		}
		if fac -= COM(n-int64(i+1), k); fac < 0 {
			fac += MOD
		}
		ans = (ans + fac*int64(a[i+1]-a[i])) % MOD
	}

	puts(ans)
}
