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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	n, k := nextInt(), nextInt()

	min, max := 0, 0
	for i := 1; i < k; i++ {
		min += i
		max += n - i + 1
	}

	ans := 0
	for i := k; i <= n+1; i++ {
		ans += max - min + 1
		ans %= mod
		min += i - 1
		max += n - i + 1
	}

	puts(ans)
}
