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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	w, h := nextInt()-1, nextInt()-1

	ans := 1
	for i := 2; i <= w+h; i++ {
		ans = (ans * i) % mod
	}
	for i := 2; i <= w; i++ {
		ans = (ans * modInv(i, mod)) % mod
	}
	for i := 2; i <= h; i++ {
		ans = (ans * modInv(i, mod)) % mod
	}
	puts(ans)
}
