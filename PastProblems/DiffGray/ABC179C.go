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

// nまでの自然数の最小の素因数
func minFactors(n int) []int {
	mf := make([]int, n+1)
	mf[0], mf[1] = -1, -1
	for i := 2; i <= n; i++ {
		for j := 1; i*j <= n; j++ {
			if mf[i*j] == 0 {
				mf[i*j] = i
			}
		}
	}
	return mf
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	mf := minFactors(n)
	ans := 0
	for c := 1; c < n; c++ {
		mp := map[int]int{}
		for d := n - c; d > 1; d /= mf[d] {
			mp[mf[d]]++
		}
		x := 1
		for _, v := range mp {
			x *= (v + 1)
		}
		ans += x
	}

	puts(ans)
}
