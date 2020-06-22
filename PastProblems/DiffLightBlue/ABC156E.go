// 解説AC

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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

const COMMAX int64 = 1000000
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

func min(nums ...int64) int64 {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	COMinit()

	n, k := int64(nextInt()), int64(nextInt())

	ans := int64(0)
	for i := int64(0); i <= min(k, n-1); i++ {
		add := (COM(n, i) * COM(n-1, n-1-i)) % MOD
		ans += add
		ans %= MOD
	}

	puts(ans)
}
