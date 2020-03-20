package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	n, m := nextInt(), nextInt()

	if abs(n-m) > 1 {
		puts(0)
		return
	}

	ans := (fact(n, mod) * fact(m, mod)) % mod
	if n == m {
		ans = (ans * 2) % mod
	}

	puts(ans)
}

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

func fact(n, mod int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret = (ret * i) % mod
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
