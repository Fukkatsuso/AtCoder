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

func getInt3() (int, int, int) {
	return getInt(), getInt(), getInt()
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m, k := getInt3()

	if n == 1 {
		puts(modPow(k, m, mod))
		return
	} else if m == 1 {
		puts(modPow(k, n, mod))
		return
	}

	ans := 0
	for x := 1; x <= k; x++ {
		// x以下の数でAを作る
		// ただし全てのAがx-1以下の場合を除く
		a := (modPow(x, n, mod) - modPow(x-1, n, mod) + mod) % mod
		// x以上の数でBを作る
		b := modPow(k-x+1, m, mod)
		ans += a * b % mod
		ans %= mod
	}

	puts(ans)
}
