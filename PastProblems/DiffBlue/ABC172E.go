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

const COMMAX int = 1000000
const MOD int = 1000000007

var fac [COMMAX]int
var finv [COMMAX]int
var inv [COMMAX]int

// テーブル作成
func COMinit() {
	fac[0] = 1
	fac[1] = 1
	finv[0] = 1
	finv[1] = 1
	inv[1] = 1
	for i := 2; i < COMMAX; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
}

// 二項係数nCk
func COM(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % MOD) % MOD
}

func perm(n, k int) int {
	return (COM(n, k) * fac[k]) % MOD
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	COMinit()
	n, m := nextInt(), nextInt()

	ans := 0
	for k := 0; k <= n; k++ {
		p1, p2 := perm(m, k), perm(m-k, n-k)
		add := (((((COM(n, k) * p1) % MOD) * p2) % MOD) * p2) % MOD
		if k%2 == 1 {
			add = -add + MOD
		}
		ans += add % MOD
		ans %= MOD
	}

	puts(ans)
}
