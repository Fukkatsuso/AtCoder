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

	n, q := nextInt(), nextInt()
	s := next()

	cnt := make([]int, n)
	for i := 1; i < n; i++ {
		cnt[i] = cnt[i-1]
		if s[i-1:i+1] == "AC" {
			cnt[i]++
		}
	}

	for i := 0; i < q; i++ {
		l, r := nextInt()-1, nextInt()-1
		puts(cnt[r] - cnt[l])
	}
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
