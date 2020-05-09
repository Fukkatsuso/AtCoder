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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

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

// 二項係数のmodを求める
func modCom(n, k, mod int) int {
	ret := 1
	for i := n - k + 1; i <= n; i++ {
		ret = (ret * i) % mod
	}
	for i := 2; i <= k; i++ {
		ret = (ret * modInv(i, mod)) % mod
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	n, k := nextInt2()

	for i := 1; i <= k; i++ {
		// k個の青いボールをi個のグループに分ける
		blue := modCom(k-1, i-1, mod)
		// n-k個の赤いボールをi+1グループに分ける
		if n-k < i-1 {
			puts(0)
			continue
		}
		red := modCom(n-k+1, i, mod)

		puts((blue * red) % mod)
	}
}
