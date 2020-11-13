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

func getInt4() (int, int, int, int) {
	return getInt(), getInt(), getInt(), getInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

const COMMAX int = 200010

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
		fac[i] = fac[i-1] * i % mod
		inv[i] = mod - inv[mod%i]*(mod/i)%mod
		finv[i] = finv[i-1] * inv[i] % mod
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
	return fac[n] * (finv[k] * finv[n-k] % mod) % mod
}

func pattern(x1, y1, x2, y2 int) int {
	dx, dy := x2-x1, y2-y1
	return COM(dx+dy, dx)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w, a, b := getInt4()

	ans := 0
	for i := b; i < w; i++ {
		// (0,0)から(h-a-1,i), (h-a,i), (h-1,w-1)と移動する
		add := pattern(0, 0, h-a-1, i) * pattern(h-a, i, h-1, w-1)
		ans += add % mod
		ans %= mod
	}

	puts(ans)
}
