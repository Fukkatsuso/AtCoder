package main

import (
	"bufio"
	"fmt"
	"os"
)

func sel(n int) int {
	return (n * (n + 1)) / 2
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()
	n := len(s)

	ans := 0
	for i := 0; i < n; {
		var l, r int
		for l, r = i, i; r+2 <= n && s[r:r+2] == "25"; r += 2 {
		}
		m := (r - l) / 2
		ans += sel(m)
		i = r + 1
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
