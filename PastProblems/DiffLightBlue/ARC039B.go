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

	n, k := nextInt(), nextInt()

	var ans int
	if n > k {
		ans = modCom(n-1+k, n-1, mod)
	} else {
		ans = modCom(n, k%n, mod)
	}

	puts(ans)
}
