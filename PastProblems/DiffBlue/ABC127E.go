// 写経AC
// 参考
// https://drken1215.hatenablog.com/entry/2019/06/15/111500

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
	mod            = 1000000007
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInt3() (int, int, int) {
	return getInt(), getInt(), getInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

const COMMAX int = 200010
const MOD int = mod

var fac [COMMAX]int
var finv [COMMAX]int
var inv [COMMAX]int

// テーブル作成
func init() {
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m, k := getInt3()

	ans := 0
	for dx := 0; dx < n; dx++ {
		for dy := 0; dy < m; dy++ {
			add := (((n - dx) * (m - dy)) % mod * (dx + dy)) % mod
			if dx != 0 && dy != 0 {
				add *= 2
			}
			add %= mod
			ans += add
			ans %= mod
		}
	}

	ans *= COM(n*m-2, k-2)
	ans %= mod

	puts(ans)
}
