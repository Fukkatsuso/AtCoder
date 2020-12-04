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

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func sum(n int, x []int) int {
	res := 0
	for i := 1; i <= n; i++ {
		add := (((2*i - n - 1 + mod) % mod) * ((x[i-1] + mod) % mod)) % mod
		res += add
		res %= mod
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	x, y := getInts(n), getInts(m)

	ans := sum(n, x) * sum(m, y)
	ans %= mod

	puts(ans)
}
