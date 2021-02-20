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
	mod            = 998244353
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
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

// modを法としたaの逆元
// 拡張ユークリッドの互除法
func modInv(a, mod int) int {
	b := mod
	u, v := 1, 0
	for b > 0 {
		t := a / b
		a -= t * b
		u -= t * v
		a, b = b, a
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w, k := getInt(), getInt(), getInt()
	t := make([][]byte, h)
	for i := range t {
		t[i] = make([]byte, w)
	}
	for i := 0; i < k; i++ {
		hi, wi, ci := getInt()-1, getInt()-1, gets()
		t[hi][wi] = ci[0]
	}

	dp := make([][]int, h+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
	}
	dp[0][0] = modPow(3, h*w-k, mod)
	r := (2 * modInv(3, mod)) % mod
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			switch t[i][j] {
			case 'R':
				dp[i][j+1] += dp[i][j]
			case 'D':
				dp[i+1][j] += dp[i][j]
			case 'X':
				dp[i][j+1] += dp[i][j]
				dp[i+1][j] += dp[i][j]
			default:
				dp[i][j+1] += dp[i][j] * r
				dp[i+1][j] += dp[i][j] * r
			}
			dp[i][j+1] %= mod
			dp[i+1][j] %= mod
		}
	}

	puts(dp[h-1][w-1])
}
